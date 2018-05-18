import { get, pick } from 'lodash';
import PropTypes from 'prop-types';
import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import { updateServiceMember, loadServiceMember } from './ducks';
import { reduxifyWizardForm } from 'shared/WizardPage/Form';
import { SwaggerField } from 'shared/JsonSchemaForm/JsonSchemaField';

const subsetOfFields = ['first_name', 'middle_name', 'last_name', 'suffix'];
const formName = 'service_member_name';
const NameWizardForm = reduxifyWizardForm(formName);

export class Name extends Component {
  componentDidMount() {
    this.props.loadServiceMember(this.props.match.params.serviceMemberId);
  }

  handleSubmit = () => {
    const pendingValues = this.props.formData.values;
    if (pendingValues) {
      const patch = pick(pendingValues, subsetOfFields);
      this.props.updateServiceMember(patch);
    }
  };

  render() {
    const {
      pages,
      pageKey,
      hasSubmitSuccess,
      error,
      currentServiceMember,
    } = this.props;
    // initialValues has to be null until there are values from the action since only the first values are taken
    const initialValues = currentServiceMember
      ? pick(currentServiceMember, subsetOfFields)
      : null;
    const serviceMemberId = this.props.match.params.serviceMemberId;
    return (
      <NameWizardForm
        handleSubmit={this.handleSubmit}
        className={formName}
        pageList={pages}
        pageKey={pageKey}
        hasSucceeded={hasSubmitSuccess}
        serverError={error}
        initialValues={initialValues}
        additionalParams={{ serviceMemberId }}
      >
        <h1 className="sm-heading">Name</h1>
        <SwaggerField
          fieldName="first_name"
          swagger={this.props.schema}
          required
        />
        <SwaggerField fieldName="middle_name" swagger={this.props.schema} />
        <SwaggerField
          fieldName="last_name"
          swagger={this.props.schema}
          required
        />
        <SwaggerField fieldName="suffix" swagger={this.props.schema} />
      </NameWizardForm>
    );
  }
}
Name.propTypes = {
  schema: PropTypes.object.isRequired,
  updateServiceMember: PropTypes.func.isRequired,
  currentServiceMember: PropTypes.object,
  error: PropTypes.object,
  hasSubmitSuccess: PropTypes.bool.isRequired,
};

function mapDispatchToProps(dispatch) {
  return bindActionCreators(
    { updateServiceMember, loadServiceMember },
    dispatch,
  );
}
function mapStateToProps(state) {
  return {
    schema: get(
      state,
      'swagger.spec.definitions.CreateServiceMemberPayload',
      {},
    ),
    formData: state.form[formName],
    ...state.serviceMember,
  };
}
export default connect(mapStateToProps, mapDispatchToProps)(Name);
