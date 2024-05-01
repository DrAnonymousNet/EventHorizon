from config import constant

notification_channel_map = {
    constant.RSVP_CONFIRMATION_TO_ATTENDEE: [constant.EMAIL],
    constant.RSVP_NOTIFICATION_TO_ORGANIZER: [
        constant.EMAIL,
        constant.PUSH_NOTIFICATION,
    ],
}
