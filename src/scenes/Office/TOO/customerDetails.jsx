import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux';
import { get, isEmpty } from 'lodash';
import {
  updateMoveTaskOrderStatus,
  getAllMoveTaskOrders,
  getMoveOrder,
  getCustomer,
  selectCustomer,
  selectMoveOrder,
} from 'shared/Entities/modules/moveTaskOrders';
import { getMTOServiceItems, selectMTOServiceItems } from 'shared/Entities/modules/mtoServiceItems';

class CustomerDetails extends Component {
  componentDidMount() {
    const { customerId, moveOrderId } = this.props.match.params;
    this.props.getCustomer(customerId);
    this.props.getMoveOrder(moveOrderId).then(({ response: { body: moveOrder } }) => {
      this.props.getAllMoveTaskOrders(moveOrder.id).then(({ response: { body: moveTaskOrder } }) => {
        moveTaskOrder.forEach(item => this.props.getMTOServiceItems(item.id));
      });
    });
  }
  render() {
    const { moveTaskOrder, customer, moveOrder, mtoServiceItems } = this.props;
    const entitlements = get(moveOrder, 'entitlement', {});
    return (
      <>
        <h1>Customer Details Page</h1>
        {customer && (
          <>
            <h2>Customer Info</h2>
            <dl>
              <dt>First Name</dt>
              <dd>{get(customer, 'first_name')}</dd>
              <dt>Last Name</dt>
              <dd>{get(customer, 'last_name')}</dd>
              <dt>ID</dt>
              <dd>{get(customer, 'id')}</dd>
              <dt>DOD ID</dt>
              <dd>{get(customer, 'dodID')}</dd>
            </dl>
          </>
        )}
        {moveOrder && (
          <>
            <h2>Move Orders</h2>
            <dt>Destination Duty Station</dt>
            <dd>{get(moveOrder, 'destinationDutyStation.name', '')}</dd>
            <dt>Destination Duty Station Address</dt>
            <dd>{JSON.stringify(get(moveOrder, 'destinationDutyStation.address', {}))} </dd>
            <dt>Origin Duty Station</dt>
            <dd>{get(moveOrder, 'originDutyStation.name', '')}</dd>
            <dt>Origin Duty Station Address</dt>
            <dd>{JSON.stringify(get(moveOrder, 'originDutyStation.address', {}))} </dd>
            {entitlements && (
              <>
                <h2>Customer Entitlements</h2>
                <dl>
                  <dt>Dependents Authorized</dt>
                  <dd>{get(entitlements, 'dependentsAuthorized', '').toString()}</dd>
                  <dt>Non Temporary Storage</dt>
                  <dd>{get(entitlements, 'nonTemporaryStorage', '').toString()}</dd>
                  <dt>Privately Owned Vehicle</dt>
                  <dd>{get(entitlements, 'privatelyOwnedVehicle', '').toString()}</dd>
                  <dt>ProGear Weight Spouse</dt>
                  <dd>{get(entitlements, 'proGearWeightSpouse')}</dd>
                  <dt>Storage In Transit</dt>
                  <dd>{get(entitlements, 'storageInTransit', '').toString()}</dd>
                  <dt>Total Dependents</dt>
                  <dd>{get(entitlements, 'totalDependents')}</dd>
                </dl>
              </>
            )}
          </>
        )}
        {!isEmpty(moveTaskOrder) && (
          <>
            <h2>Move Task Order</h2>
            <dl>
              <dt>ID</dt>
              <dd>{get(moveTaskOrder, 'id')}</dd>
              <dt>Is Available to Prime</dt>
              <dd>{get(moveTaskOrder, 'isAvailableToPrime').toString()}</dd>
              <dt>Is Canceled</dt>
              <dd>{get(moveTaskOrder, 'isCanceled', false).toString()}</dd>
            </dl>

            <h2>MTO Service Items</h2>
            <table>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Move Task Order ID</th>
                  <th>Rate Engine Service ID</th>
                  <th>Rate Engine Service Code</th>
                  <th>Rate Engine Service Name</th>
                  <th>MTO Shipment ID</th>
                </tr>
              </thead>
              <tbody>
                {mtoServiceItems.map(items => {
                  return (
                    <Fragment key={items.id}>
                      <tr>
                        <td>{items.id}</td>
                        <td>{items.moveTaskOrderID}</td>
                        <td>{items.reServiceID}</td>
                        <td>{items.reServiceCode}</td>
                        <td>{items.reServiceName}</td>
                        <td>{items.mtoShipmentID}</td>
                      </tr>
                    </Fragment>
                  );
                })}
              </tbody>
            </table>

            <div>
              <button onClick={() => this.props.updateMoveTaskOrderStatus(moveTaskOrder.id)}>Send to Prime</button>
            </div>
          </>
        )}
      </>
    );
  }
}

const mapStateToProps = (state, ownProps) => {
  const moveOrder = selectMoveOrder(state, ownProps.match.params.moveOrderId);
  const mto = Object.values(get(state, 'entities.moveTaskOrder', {}))[0];
  return {
    moveOrder,
    customer: selectCustomer(state, ownProps.match.params.customerId),

    // TODO: Change when we start making use of multiple move task orders
    mtoServiceItems: selectMTOServiceItems(state, new Set([mto?.id])),
    moveTaskOrder: mto,
  };
};

const mapDispatchToProps = {
  getMoveOrder,
  getAllMoveTaskOrders,
  updateMoveTaskOrderStatus,
  getCustomer,
  getMTOServiceItems,
};

export default connect(mapStateToProps, mapDispatchToProps)(CustomerDetails);
