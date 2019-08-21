import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { capitalize, get, includes } from 'lodash';

import { NavTab, RoutedTabs } from 'react-router-tabs';
import { Link, NavLink, Redirect, Switch } from 'react-router-dom';
import LoadingPlaceholder from 'shared/LoadingPlaceholder';
import PrivateRoute from 'shared/User/PrivateRoute';
import Alert from 'shared/Alert'; // eslint-disable-line
import ToolTip from 'shared/ToolTip';
import ComboButton from 'shared/ComboButton';
import { DropDown, DropDownItem } from 'shared/ComboButton/dropdown';
import DocumentList from 'shared/DocumentViewer/DocumentList';
import AccountingPanel from './AccountingPanel';
import BackupInfoPanel from './BackupInfoPanel';
import CustomerInfoPanel from './CustomerInfoPanel';
import OrdersPanel from './OrdersPanel';
import PaymentsPanel from './Ppm/PaymentsPanel';
import DatesAndLocationPanel from './Ppm/DatesAndLocationsPanel';
import PPMEstimatesPanel from './Ppm/PPMEstimatesPanel';
import StoragePanel from './Ppm/StoragePanel';
import ExpensesPanel from './Ppm/ExpensesPanel';
import NetWeightPanel from './Ppm/NetWeightPanel';
import { withContext } from 'shared/AppContext';
import ConfirmWithReasonButton from 'shared/ConfirmWithReasonButton';

import { getRequestStatus } from 'shared/Swagger/selectors';
import { resetRequests } from 'shared/Swagger/request';
import { getStorageInTransitsForShipment, selectStorageInTransits } from 'shared/Entities/modules/storageInTransits';
import { getAllTariff400ngItems, selectTariff400ngItems } from 'shared/Entities/modules/tariff400ngItems';
import {
  selectSortedShipmentLineItems,
  fetchAndCalculateShipmentLineItems,
} from 'shared/Entities/modules/shipmentLineItems';
import { getAllInvoices } from 'shared/Entities/modules/invoices';
import { approvePPM, loadPPMs, selectPPMForMove, selectReimbursement } from 'shared/Entities/modules/ppms';
import { loadBackupContacts, loadServiceMember, selectServiceMember } from 'shared/Entities/modules/serviceMembers';
import { loadOrders, loadOrdersLabel, selectOrders } from 'shared/Entities/modules/orders';
import {
  approveShipment,
  getPublicShipment,
  selectShipment,
  selectShipmentStatus,
  updatePublicShipment,
} from 'shared/Entities/modules/shipments';
import { getTspForShipment } from 'shared/Entities/modules/transportationServiceProviders';
import { getServiceAgentsForShipment, selectServiceAgentsForShipment } from 'shared/Entities/modules/serviceAgents';

import { showBanner, removeBanner } from './ducks';
import {
  loadMove,
  loadMoveLabel,
  selectMove,
  selectMoveStatus,
  approveBasics,
  cancelMove,
} from 'shared/Entities/modules/moves';
import { formatDate } from 'shared/formatters';
import { getMoveDocumentsForMove, selectAllDocumentsForMove } from 'shared/Entities/modules/moveDocuments';

import FontAwesomeIcon from '@fortawesome/react-fontawesome';
import faPhone from '@fortawesome/fontawesome-free-solid/faPhone';
import faEmail from '@fortawesome/fontawesome-free-solid/faEnvelope';
import faClock from '@fortawesome/fontawesome-free-solid/faClock';
import faCheck from '@fortawesome/fontawesome-free-solid/faCheck';
import faExclamationCircle from '@fortawesome/fontawesome-free-solid/faExclamationCircle';
import faPlayCircle from '@fortawesome/fontawesome-free-solid/faPlayCircle';
import moment from 'moment';

const BasicsTabContent = props => {
  return (
    <div className="office-tab">
      <OrdersPanel title="Orders" moveId={props.moveId} />
      <CustomerInfoPanel title="Customer Info" serviceMember={props.serviceMember} />
      <BackupInfoPanel title="Backup Info" serviceMember={props.serviceMember} />
      <AccountingPanel title="Accounting" serviceMember={props.serviceMember} moveId={props.moveId} />
    </div>
  );
};

const PPMTabContent = props => {
  return (
    <div className="office-tab">
      <PaymentsPanel title="Payments" moveId={props.moveId} />
      {props.ppmPaymentRequested && (
        <>
          <ExpensesPanel title="Expenses" moveId={props.moveId} moveDocuments={props.moveDocuments} />
          <StoragePanel title="Storage" moveId={props.moveId} moveDocuments={props.moveDocuments} />
          <DatesAndLocationPanel title="Dates & Locations" moveId={props.moveId} />
          <NetWeightPanel
            title="Weights"
            moveId={props.moveId}
            ppmPaymentRequestedFlag={props.ppmPaymentRequestedFlag}
          />
        </>
      )}

      <PPMEstimatesPanel title="Estimates" moveId={props.moveId} />
    </div>
  );
};

const ReferrerQueueLink = props => {
  const pathname = props.history.location.state ? props.history.location.state.referrerPathname : '';
  switch (pathname) {
    case '/queues/ppm':
      return (
        <NavLink to="/queues/ppm" activeClassName="usa-current">
          <span>All PPMs Queue</span>
        </NavLink>
      );
    case '/queues/ppm_payment_requested':
      return (
        <NavLink to="/queues/ppm_payment_requested" activeClassName="usa-current">
          <span>Payment Requests PPMs Queue</span>
        </NavLink>
      );
    case '/queues/all':
      return (
        <NavLink to="/queues/all" activeClassName="usa-current">
          <span>All Moves Queue</span>
        </NavLink>
      );
    default:
      return (
        <NavLink to="/queues/new" activeClassName="usa-current">
          <span>New Moves/Shipments Queue</span>
        </NavLink>
      );
  }
};

class MoveInfo extends Component {
  state = {
    redirectToHome: false,
  };

  componentDidMount() {
    const { moveId } = this.props;
    this.props.loadMove(moveId);
    this.props.getMoveDocumentsForMove(moveId);
    this.props.getAllTariff400ngItems(true);
    this.props.loadPPMs(moveId);
  }

  componentDidUpdate(prevProps) {
    const {
      loadBackupContacts,
      loadOrders,
      loadMoveIsSuccess,
      loadServiceMember,
      ordersId,
      serviceMemberId,
      shipmentId,
    } = this.props;
    if (loadMoveIsSuccess !== prevProps.loadMoveIsSuccess && loadMoveIsSuccess) {
      loadOrders(ordersId);
      loadServiceMember(serviceMemberId);
      loadBackupContacts(serviceMemberId);
      if (shipmentId) {
        this.getAllShipmentInfo(shipmentId);
      }
    }
  }

  componentWillUnmount() {
    this.props.resetRequests();
  }

  get allAreApproved() {
    const { moveStatus, ppm } = this.props;
    const moveApproved = moveStatus === 'APPROVED';
    const ppmApproved = includes(['APPROVED', 'PAYMENT_REQUESTED', 'COMPLETED'], ppm.status);
    return moveApproved && ppmApproved;
  }
  getAllShipmentInfo = shipmentId => {
    this.props.getTspForShipment(shipmentId);
    this.props.getPublicShipment(shipmentId);
    this.props.fetchAndCalculateShipmentLineItems(shipmentId, this.props.shipmentStatus);
    this.props.getAllInvoices(shipmentId);
    this.props.getServiceAgentsForShipment(shipmentId);
    this.props.getStorageInTransitsForShipment(shipmentId);
  };

  approveBasics = () => {
    this.props.approveBasics(this.props.moveId);
  };

  approvePPM = () => {
    const approveDate = moment().format();
    this.props.approvePPM(this.props.ppm.id, approveDate);
  };

  approveShipment = () => {
    const approveDate = moment().format();
    this.props.approveShipment(this.props.shipmentId, approveDate);
  };

  cancelMoveAndRedirect = cancelReason => {
    const messageLines = [
      `Move #${this.props.move.locator} for ${this.props.serviceMember.last_name}, ${this.props.serviceMember.first_name} has been canceled`,
      'An email confirmation has been sent to the customer.',
    ];
    this.props.cancelMove(this.props.moveId, cancelReason).then(() => {
      this.props.showBanner({ messageLines });
      setTimeout(() => this.props.removeBanner(), 10000);
      this.setState({ redirectToHome: true });
    });
  };

  renderPPMTabStatus = () => {
    if (this.props.ppm.status === 'APPROVED') {
      if (this.props.ppmAdvance.status === 'APPROVED' || !this.props.ppmAdvance.status) {
        return (
          <span className="status">
            <FontAwesomeIcon className="icon approval-ready" icon={faCheck} />
            Move pending
          </span>
        );
      } else {
        return (
          <span className="status">
            <FontAwesomeIcon className="icon approval-waiting" icon={faClock} />
            Payment Requested
          </span>
        );
      }
    } else {
      return (
        <span className="status">
          <FontAwesomeIcon className="icon approval-waiting" icon={faClock} />
          In review
        </span>
      );
    }
  };

  render() {
    const {
      move,
      moveDocuments,
      moveStatus,
      orders,
      ppm,
      shipment,
      shipmentStatus,
      serviceMember,
      upload,
    } = this.props;
    const isPPM = move.selected_move_type === 'PPM';
    const showDocumentViewer = this.props.context.flags.documentViewer;
    const moveInfoComboButton = this.props.context.flags.moveInfoComboButton;
    const ordersComplete = Boolean(
      orders.orders_number && orders.orders_type_detail && orders.department_indicator && orders.tac && orders.sac,
    );
    const ppmPaymentRequested = includes(['PAYMENT_REQUESTED', 'COMPLETED'], ppm.status);
    const ppmApproved = includes(['APPROVED', 'PAYMENT_REQUESTED', 'COMPLETED'], ppm.status);
    const moveApproved = moveStatus === 'APPROVED';
    const hhgCantBeCanceled = includes(['IN_TRANSIT', 'DELIVERED'], shipmentStatus);

    const moveDate = isPPM ? ppm.original_move_date : shipment && shipment.requested_pickup_date;

    const uploadDocumentUrl = `/moves/${this.props.moveId}/documents/new`;

    if (this.state.redirectToHome) {
      return <Redirect to="/" />;
    }

    if (!this.props.loadDependenciesHasSuccess && !this.props.loadDependenciesHasError) return <LoadingPlaceholder />;
    if (this.props.loadDependenciesHasError)
      return (
        <div className="usa-grid">
          <div className="usa-width-one-whole error-message">
            <Alert type="error" heading="An error occurred">
              Something went wrong contacting the server.
            </Alert>
          </div>
        </div>
      );

    return (
      <div>
        <div className="usa-grid grid-wide">
          <div className="usa-width-two-thirds">
            <h1>
              Move Info: {serviceMember.last_name}, {serviceMember.first_name}
            </h1>
          </div>
          <div className="usa-width-one-third nav-controls">
            <ReferrerQueueLink history={this.props.history} />
          </div>
        </div>
        <div className="usa-grid grid-wide">
          <div className="usa-width-one-whole">
            <ul className="move-info-header-meta">
              <li>ID# {serviceMember.edipi}&nbsp;</li>
              <li>
                {serviceMember.telephone}
                {serviceMember.phone_is_preferred && (
                  <FontAwesomeIcon className="icon icon-grey" icon={faPhone} flip="horizontal" />
                )}
                {serviceMember.email_is_preferred && <FontAwesomeIcon className="icon icon-grey" icon={faEmail} />}
                &nbsp;
              </li>
              <li>Locator# {move.locator}&nbsp;</li>
              {shipment.gbl_number && <li>GBL# {shipment.gbl_number}&nbsp;</li>}
              <li>Move date {formatDate(moveDate)}&nbsp;</li>
            </ul>
          </div>
        </div>

        <div className="usa-grid grid-wide tabs">
          <div className="usa-width-three-fourths">
            <RoutedTabs startPathWith={this.props.match.url}>
              <NavTab to="/basics">
                <span className="title" data-cy="basics-tab">
                  Basics
                </span>
                <span className="status">
                  <FontAwesomeIcon className="icon" icon={faPlayCircle} />
                  {capitalize(this.props.moveStatus)}
                </span>
              </NavTab>
              {isPPM && (
                <NavTab to="/ppm">
                  <span className="title" data-cy="ppm-tab">
                    PPM
                  </span>
                  {this.renderPPMTabStatus()}
                </NavTab>
              )}
            </RoutedTabs>

            <div className="tab-content">
              <Switch>
                <PrivateRoute
                  exact
                  path={`${this.props.match.url}`}
                  render={() => (
                    <Redirect
                      replace
                      to={{ pathname: `${this.props.match.url}/basics`, state: this.props.history.location.state }}
                    />
                  )}
                />
                <PrivateRoute path={`${this.props.match.path}/basics`}>
                  <BasicsTabContent moveId={this.props.moveId} serviceMember={this.props.serviceMember} />
                </PrivateRoute>
                <PrivateRoute path={`${this.props.match.path}/ppm`}>
                  <PPMTabContent
                    ppmPaymentRequestedFlag={this.props.context.flags.ppmPaymentRequest}
                    moveId={this.props.moveId}
                    ppmPaymentRequested={ppmPaymentRequested}
                    moveDocuments={moveDocuments}
                  />
                </PrivateRoute>
              </Switch>
            </div>
          </div>
          <div className="usa-width-one-fourth">
            <div>
              {this.props.approveMoveHasError && (
                <Alert type="warning" heading="Unable to approve">
                  Please fill out missing data
                </Alert>
              )}
              <div>
                <ToolTip
                  disabled={ordersComplete}
                  textStyle="tooltiptext-large"
                  toolTipText="Some information about the move is missing or contains errors. Please fix these problems before approving."
                >
                  {moveInfoComboButton && (
                    <ComboButton
                      allAreApproved={this.allAreApproved}
                      buttonText={`Approve${this.allAreApproved ? 'd' : ''}`}
                      disabled={this.allAreApproved || !ordersComplete}
                    >
                      <DropDown>
                        <DropDownItem
                          value="Approve Basics"
                          disabled={moveApproved || !ordersComplete}
                          onClick={this.approveBasics}
                        />
                        {isPPM && (
                          <DropDownItem
                            disabled={ppmApproved || !moveApproved || !ordersComplete}
                            onClick={this.approvePPM}
                            value="Approve PPM"
                          />
                        )}
                      </DropDown>
                    </ComboButton>
                  )}
                </ToolTip>
                <ConfirmWithReasonButton
                  buttonTitle="Cancel Move"
                  reasonPrompt="Why is the move being canceled?"
                  warningPrompt="Are you sure you want to cancel the entire move?"
                  onConfirm={this.cancelMoveAndRedirect}
                  buttonDisabled={hhgCantBeCanceled}
                />
              </div>
            </div>
            <div className="documents">
              <h2 className="extras usa-heading">Documents</h2>
              {!upload ? (
                <p>No orders have been uploaded.</p>
              ) : (
                <div>
                  {moveApproved ? (
                    <div className="panel-field">
                      <FontAwesomeIcon style={{ color: 'green' }} className="icon" icon={faCheck} />
                      <Link to={`/moves/${move.id}/orders`} target="_blank">
                        Orders ({formatDate(upload.created_at)})
                      </Link>
                    </div>
                  ) : (
                    <div className="panel-field">
                      <FontAwesomeIcon style={{ color: 'red' }} className="icon" icon={faExclamationCircle} />
                      <Link to={`/moves/${move.id}/orders`} target="_blank">
                        Orders ({formatDate(upload.created_at)})
                      </Link>
                    </div>
                  )}
                </div>
              )}
              {showDocumentViewer && (
                <DocumentList
                  detailUrlPrefix={`/moves/${this.props.moveId}/documents`}
                  moveDocuments={moveDocuments}
                  uploadDocumentUrl={uploadDocumentUrl}
                />
              )}
            </div>
          </div>
        </div>
      </div>
    );
  }
}

MoveInfo.defaultProps = {
  move: {},
};

MoveInfo.propTypes = {
  loadMove: PropTypes.func.isRequired,
  context: PropTypes.shape({
    flags: PropTypes.shape({
      documentViewer: PropTypes.bool,
      sitPanel: PropTypes.bool,
    }).isRequired,
  }).isRequired,
};

const mapStateToProps = (state, ownProps) => {
  const moveId = ownProps.match.params.moveId;
  const move = selectMove(state, moveId);
  const shipmentId = get(move, 'shipments.0.id');
  const ppm = selectPPMForMove(state, moveId);
  const ordersId = move.orders_id;
  const orders = selectOrders(state, ordersId);
  const serviceMemberId = move.service_member_id;
  const serviceMember = selectServiceMember(state, serviceMemberId);
  const loadOrdersStatus = getRequestStatus(state, loadOrdersLabel);
  const loadMoveIsSuccess = getRequestStatus(state, loadMoveLabel).isSuccess;

  return {
    approveMoveHasError: get(state, 'office.moveHasApproveError'),
    errorMessage: get(state, 'office.error'),
    loadDependenciesHasError: loadOrdersStatus.error,
    loadDependenciesHasSuccess: loadOrdersStatus.isSuccess,
    loadMoveIsSuccess,
    moveDocuments: selectAllDocumentsForMove(state, moveId),
    ppm,
    move,
    moveId,
    moveStatus: selectMoveStatus(state, moveId),
    orders,
    ordersId,
    officeShipment: get(state, 'office.officeShipment', {}),
    ppmAdvance: selectReimbursement(state, ppm.advance),
    serviceAgents: selectServiceAgentsForShipment(state, shipmentId),
    serviceMember,
    serviceMemberId,
    shipment: selectShipment(state, shipmentId),
    shipmentId,
    shipmentLineItems: selectSortedShipmentLineItems(state),
    shipmentStatus: selectShipmentStatus(state, shipmentId),
    storageInTransits: selectStorageInTransits(state, shipmentId),
    swaggerError: get(state, 'swagger.hasErrored'),
    tariff400ngItems: selectTariff400ngItems(state),
    upload: get(orders, 'uploaded_orders.uploads.0', {}),
  };
};

const mapDispatchToProps = dispatch =>
  bindActionCreators(
    {
      getPublicShipment,
      updatePublicShipment,
      getMoveDocumentsForMove,
      approveBasics,
      approvePPM,
      approveShipment,
      cancelMove,
      getAllTariff400ngItems,
      fetchAndCalculateShipmentLineItems,
      getAllInvoices,
      getTspForShipment,
      getServiceAgentsForShipment,
      getStorageInTransitsForShipment,
      showBanner,
      removeBanner,
      loadMove,
      loadPPMs,
      loadServiceMember,
      loadBackupContacts,
      loadOrders,
      resetRequests,
    },
    dispatch,
  );

const connectedMoveInfo = withContext(
  connect(
    mapStateToProps,
    mapDispatchToProps,
  )(MoveInfo),
);
export { connectedMoveInfo as default, ReferrerQueueLink };
