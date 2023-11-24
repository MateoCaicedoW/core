package forms

import (
	"html/template"

	"github.com/gobuffalo/tags/v3"
	"github.com/gobuffalo/tags/v3/form"
	"github.com/leapkit/core/render/hctx"
)

// FormFor implements a `github.com/gobuffalo/plush` helper around the
// form.NewFormFor function in the `github.com/gobuffalo/tags/form` package
func FormFor(model interface{}, opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if opts == nil {
		opts = tags.Options{}
	}
	return helper(opts, help, func(opts tags.Options) helperable {
		return form.NewFormFor(model, opts)
	})
}

// RemoteFormFor implements a `github.com/gobuffalo/plush` helper around the
// form.NewFormFor function in the `github.com/gobuffalo/tags/form` package
//
//	<%= remoteFormFor(widget, {method: "POST"}) { %>
//	<% } %>
//	<form action="/widgets/b6b0ab24-19ae-4cdd-ad73-c5ecbddd6f91" id="widget-form" method="POST"><input name="_method" type="hidden" value="PUT"></form>
func RemoteFormFor(model interface{}, opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if opts == nil {
		opts = tags.Options{}
	}
	opts["data-remote"] = true

	return FormFor(model, opts, help)
}