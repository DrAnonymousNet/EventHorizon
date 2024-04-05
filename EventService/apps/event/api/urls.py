from django.urls import include, path

urlpatterns = [
    path(
        "",
        include(("apps.event.api.v1.urls", ""), namespace="event"),
    ),
]
