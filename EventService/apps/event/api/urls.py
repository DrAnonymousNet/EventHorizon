from django.urls import include, path


urlpatterns = [
    path("", include("apps.event.api.v1.urls", namespace="event"))
    # path(
    #     "",
    #     include("apps.event.api.v1.urls", "", namespace="event", app_name="event"),
    # ),
]
