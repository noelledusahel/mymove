import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import PropTypes from 'prop-types';
import { get } from 'lodash';
import { connect } from 'react-redux';

import { selectMoveDocument } from 'shared/Entities/modules/moveDocuments';
import DocumentContent from './DocumentContent';

// export const DocumentUploadViewer = ({ moveDocument }) => {
//   const uploadModels = get(moveDocument, 'document.uploads', []);
//   return uploadModels.map(({ url, filename, content_type }) => (
//     <DocumentContent key={url} url={url} filename={filename} contentType={content_type} />
//   ));
// };

class DocumentUploadViewer extends Component {
  render() {
    const { moveDocument } = this.props;
    const uploadModels = get(moveDocument, 'document.uploads', []);
    return uploadModels.map(({ url, filename, content_type, tags }) => (
      <DocumentContent key={url} url={url} filename={filename} contentType={content_type} tags={tags} />
    ));
  }
}

const { shape, string, number, arrayOf } = PropTypes;

DocumentUploadViewer.propTypes = {
  moveDocument: shape({
    document: shape({
      id: string.isRequired,
      service_member_id: string.isRequired,
      uploads: arrayOf(
        shape({
          byes: number,
          content_type: string.isRequired,
          created_at: string.isRequired,
          filename: string.isRequired,
          id: string.isRequired,
          update_at: string,
          url: string.isRequired,
        }),
      ).isRequired,
    }),
    id: string.isRequired,
    move_document_type: string.isRequired,
    move_id: string.isRequired,
    notes: string,
    personally_procured_move_id: string,
    status: string.isRequired,
    title: string.isRequired,
  }).isRequired,
};

function mapStateToProps(state, props) {
  const moveDocumentId = props.match.params.moveDocumentId;
  return {
    moveDocument: selectMoveDocument(state, moveDocumentId),
  };
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators({}, dispatch);
}
export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(DocumentUploadViewer);
