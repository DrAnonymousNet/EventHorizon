from rest_framework.viewsets import ModelViewSet
from apps.event.api.v1.serializers import EventAttendanceSerializer, EventFullSerializer, EventMinimalSerializer
from apps.event.models import Event, EventAttendance
from rest_framework.permissions import IsAuthenticatedOrReadOnly, SAFE_METHODS
from rest_framework.decorators import action
from rest_framework.response import Response
from rest_framework import status
from drf_yasg.utils import swagger_auto_schema

class EventAPIViewSet(ModelViewSet):
    serializer_class = EventFullSerializer
    queryset = Event.objects.all()
    permission_classes = IsAuthenticatedOrReadOnly
    serializer_class = EventFullSerializer
    lookup_field = "uid"

    def get_serializer_class(self):
        if self.request.query_params.get("full", True) or self.request.method not in SAFE_METHODS:
            return super().get_serializer_class()
        return EventMinimalSerializer
    
    
    @swagger_auto_schema(
        
        responses={200: {
            "msg":"You have rsvp successfuly",
            "data":EventAttendanceSerializer()
        },
        400:{
            "msg":"validation failed",
            "data":[]
        }
        },
    )
    @action(methods=["post"], url_path="rsvp")
    def rsvp(self, request, *args, **kwargs):
        serializer = EventAttendanceSerializer(data=request.data)
        if not serializer.is_valid():
            return Response(status=status.HTTP_400_BAD_REQUEST,
                data= {
                    "msg":"validation failed",
                    "data":serializer.errors
                }
            )
        serializer.save()
        return Response(status=status.HTTP_200_OK, data={
            "msg":"You have rsvp successfuly",
            "data":serializer.data
        })
    

    

        
