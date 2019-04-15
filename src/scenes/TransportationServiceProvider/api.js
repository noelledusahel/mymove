import { getPublicClient, checkResponse } from 'shared/Swagger/api';
import { formatPayload } from 'shared/utils';

// SHIPMENT QUEUE
export async function RetrieveShipmentsForTSP(queueType) {
  const queueToStatus = {
    new: ['AWARDED'],
    accepted: ['ACCEPTED'],
    approved: ['APPROVED'],
    in_transit: ['IN_TRANSIT'],
    delivered: ['DELIVERED'],
    all: [],
  };
  /* eslint-disable security/detect-object-injection */
  const status = (queueType && queueToStatus[queueType] && queueToStatus[queueType].join(',')) || '';
  /* eslint-enable security/detect-object-injection */
  const client = await getPublicClient();
  const response = await client.apis.shipments.indexShipments({
    status,
  });
  checkResponse(response, 'failed to retrieve moves due to server error');
  return response.body;
}

// SHIPMENT ACCEPT
export async function AcceptShipment(shipmentId) {
  const client = await getPublicClient();
  const response = await client.apis.shipments.acceptShipment({
    shipmentId: shipmentId,
  });
  checkResponse(response, 'failed to accept shipment due to server error');
  return response.body;
}

export async function TransportShipment(shipmentId, payload) {
  const client = await getPublicClient();
  const payloadDef = client.spec.definitions.TransportPayload;
  const response = await client.apis.shipments.transportShipment({
    shipmentId,
    payload: formatPayload(payload, payloadDef),
  });
  checkResponse(response, 'failed to pick up shipment due to server error');
  return response.body;
}

export async function DeliverShipment(shipmentId, payload) {
  const client = await getPublicClient();
  const payloadDef = client.spec.definitions.ActualDeliveryDate;
  const response = await client.apis.shipments.deliverShipment({
    shipmentId,
    payload: formatPayload(payload, payloadDef),
  });
  checkResponse(response, 'failed to deliver shipment due to server error');
  return response.body;
}

export async function CompletePmSurvey(shipmentId) {
  const client = await getPublicClient();
  const response = await client.apis.shipments.completePmSurvey({
    shipmentId,
  });
  checkResponse(response, 'failed to complete pre-move survey due to server error');
  return response.body;
}

// ServiceAgents
export async function IndexServiceAgents(shipmentId) {
  const client = await getPublicClient();
  const response = await client.apis.service_agents.indexServiceAgents({
    shipmentId,
  });
  checkResponse(response, 'failed to load service agents due to server error');
  return response.body;
}

// All documents for shipment
export async function GetAllShipmentDocuments(shipmentId) {
  const client = await getPublicClient();
  const response = await client.apis.move_docs.indexMoveDocuments({
    shipmentId,
  });
  checkResponse(response, 'failed to load shipment documents due to server error');
  return response.body;
}
