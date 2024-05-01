def str_to_bool(value):
    return value.lower() in (
        "yes",
        "true",
        "t",
        "1",
    )  # Returns True if value is one of these
