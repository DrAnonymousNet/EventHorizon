import grpc
from django.conf import settings

from . import notifier_pb2, notifier_pb2_grpc


def notification_service_client(
    event_uid, message, subject, recipient, message_type, notification_channels
):
    with grpc.insecure_channel(settings.NOTIFICATION_SERVICE_ENDPOINT) as channel:
        stub = notifier_pb2_grpc.NotifierServiceStub(channel)
        response = stub.Notify(
            notifier_pb2.NotifyRequest(
                message=message,
                subject=subject,
                recipients=recipient,
                message_type=message_type,
                notification_channels=notification_channels,
            )
        )
        print("Notifier client received: ", response, response.message)
