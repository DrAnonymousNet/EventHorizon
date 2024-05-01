from rest_framework import status, viewsets
from rest_framework.response import Response


class ResponseViewMixin:
    def dispatch(self, request, *args, **kwargs):
        # Call the parent dispatch to get the typical response
        self.response = super().dispatch(request, *args, **kwargs)

        # Return early if there's already a "msg" key in the response data
        if "msg" in self.response:
            return self.response

        # Getting the model name for the messages
        model_name = self.queryset.model._meta.verbose_name.title()

        # Checking status codes and customizing messages
        if self.response.status_code == status.HTTP_404_NOT_FOUND:
            self.response.data["msg"] = f"{model_name} not found."
        elif self.response.status_code == status.HTTP_403_FORBIDDEN:
            self.response.data["msg"] = "Permission denied."

        # Handling for CRUD operations based on the action attribute
        if self.action == "create":
            if self.response.status_code == status.HTTP_201_CREATED:
                self.response.data["msg"] = f"{model_name} created successfully."
        elif self.action == "retrieve":
            if self.response.status_code == status.HTTP_200_OK:
                self.response.data["msg"] = f"{model_name} retrieved successfully."
        elif self.action == "update":
            if self.response.status_code == status.HTTP_200_OK:
                self.response.data["msg"] = f"{model_name} updated successfully."
        elif self.action == "partial_update":
            if self.response.status_code == status.HTTP_200_OK:
                self.response.data[
                    "msg"
                ] = f"{model_name} partially updated successfully."
        elif self.action == "destroy":
            if self.response.status_code == status.HTTP_204_NO_CONTENT:
                self.response.data["msg"] = f"{model_name} deleted successfully."

        # Return the modified response
        return self.response


# Example of how to use ResponseViewMixin in a DRF ViewSet
# class ExampleViewSet(ResponseViewMixin, viewsets.ModelViewSet):
#     queryset = ExampleModel.objects.all()
#     serializer_class = ExampleSerializer
