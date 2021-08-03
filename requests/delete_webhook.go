package requests

type DeleteWebhook struct {
	DropPendingUpdates bool
}

func (r *DeleteWebhook) IsMultipart() bool {
	return false
}

func (r *DeleteWebhook) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.DropPendingUpdates {
		values["drop_pending_updates"] = "1"
	}

	return
}
