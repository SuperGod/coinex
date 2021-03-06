// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUserUpdateParams creates a new UserUpdateParams object
// with the default values initialized.
func NewUserUpdateParams() *UserUpdateParams {
	var ()
	return &UserUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserUpdateParamsWithTimeout creates a new UserUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserUpdateParamsWithTimeout(timeout time.Duration) *UserUpdateParams {
	var ()
	return &UserUpdateParams{

		timeout: timeout,
	}
}

// NewUserUpdateParamsWithContext creates a new UserUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserUpdateParamsWithContext(ctx context.Context) *UserUpdateParams {
	var ()
	return &UserUpdateParams{

		Context: ctx,
	}
}

// NewUserUpdateParamsWithHTTPClient creates a new UserUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserUpdateParamsWithHTTPClient(client *http.Client) *UserUpdateParams {
	var ()
	return &UserUpdateParams{
		HTTPClient: client,
	}
}

/*UserUpdateParams contains all the parameters to send to the API endpoint
for the user update operation typically these are written to a http.Request
*/
type UserUpdateParams struct {

	/*Country
	  Country of residence.

	*/
	Country *string
	/*Firstname*/
	Firstname *string
	/*Lastname*/
	Lastname *string
	/*NewPassword*/
	NewPassword *string
	/*NewPasswordConfirm*/
	NewPasswordConfirm *string
	/*OldPassword*/
	OldPassword *string
	/*PgpPubKey
	  PGP Public Key. If specified, automated emails will be sentwith this key.

	*/
	PgpPubKey *string
	/*Username
	  Username can only be set once. To reset, email support.

	*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user update params
func (o *UserUpdateParams) WithTimeout(timeout time.Duration) *UserUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user update params
func (o *UserUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user update params
func (o *UserUpdateParams) WithContext(ctx context.Context) *UserUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user update params
func (o *UserUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user update params
func (o *UserUpdateParams) WithHTTPClient(client *http.Client) *UserUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user update params
func (o *UserUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCountry adds the country to the user update params
func (o *UserUpdateParams) WithCountry(country *string) *UserUpdateParams {
	o.SetCountry(country)
	return o
}

// SetCountry adds the country to the user update params
func (o *UserUpdateParams) SetCountry(country *string) {
	o.Country = country
}

// WithFirstname adds the firstname to the user update params
func (o *UserUpdateParams) WithFirstname(firstname *string) *UserUpdateParams {
	o.SetFirstname(firstname)
	return o
}

// SetFirstname adds the firstname to the user update params
func (o *UserUpdateParams) SetFirstname(firstname *string) {
	o.Firstname = firstname
}

// WithLastname adds the lastname to the user update params
func (o *UserUpdateParams) WithLastname(lastname *string) *UserUpdateParams {
	o.SetLastname(lastname)
	return o
}

// SetLastname adds the lastname to the user update params
func (o *UserUpdateParams) SetLastname(lastname *string) {
	o.Lastname = lastname
}

// WithNewPassword adds the newPassword to the user update params
func (o *UserUpdateParams) WithNewPassword(newPassword *string) *UserUpdateParams {
	o.SetNewPassword(newPassword)
	return o
}

// SetNewPassword adds the newPassword to the user update params
func (o *UserUpdateParams) SetNewPassword(newPassword *string) {
	o.NewPassword = newPassword
}

// WithNewPasswordConfirm adds the newPasswordConfirm to the user update params
func (o *UserUpdateParams) WithNewPasswordConfirm(newPasswordConfirm *string) *UserUpdateParams {
	o.SetNewPasswordConfirm(newPasswordConfirm)
	return o
}

// SetNewPasswordConfirm adds the newPasswordConfirm to the user update params
func (o *UserUpdateParams) SetNewPasswordConfirm(newPasswordConfirm *string) {
	o.NewPasswordConfirm = newPasswordConfirm
}

// WithOldPassword adds the oldPassword to the user update params
func (o *UserUpdateParams) WithOldPassword(oldPassword *string) *UserUpdateParams {
	o.SetOldPassword(oldPassword)
	return o
}

// SetOldPassword adds the oldPassword to the user update params
func (o *UserUpdateParams) SetOldPassword(oldPassword *string) {
	o.OldPassword = oldPassword
}

// WithPgpPubKey adds the pgpPubKey to the user update params
func (o *UserUpdateParams) WithPgpPubKey(pgpPubKey *string) *UserUpdateParams {
	o.SetPgpPubKey(pgpPubKey)
	return o
}

// SetPgpPubKey adds the pgpPubKey to the user update params
func (o *UserUpdateParams) SetPgpPubKey(pgpPubKey *string) {
	o.PgpPubKey = pgpPubKey
}

// WithUsername adds the username to the user update params
func (o *UserUpdateParams) WithUsername(username *string) *UserUpdateParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the user update params
func (o *UserUpdateParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *UserUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Country != nil {

		// form param country
		var frCountry string
		if o.Country != nil {
			frCountry = *o.Country
		}
		fCountry := frCountry
		if fCountry != "" {
			if err := r.SetFormParam("country", fCountry); err != nil {
				return err
			}
		}

	}

	if o.Firstname != nil {

		// form param firstname
		var frFirstname string
		if o.Firstname != nil {
			frFirstname = *o.Firstname
		}
		fFirstname := frFirstname
		if fFirstname != "" {
			if err := r.SetFormParam("firstname", fFirstname); err != nil {
				return err
			}
		}

	}

	if o.Lastname != nil {

		// form param lastname
		var frLastname string
		if o.Lastname != nil {
			frLastname = *o.Lastname
		}
		fLastname := frLastname
		if fLastname != "" {
			if err := r.SetFormParam("lastname", fLastname); err != nil {
				return err
			}
		}

	}

	if o.NewPassword != nil {

		// form param newPassword
		var frNewPassword string
		if o.NewPassword != nil {
			frNewPassword = *o.NewPassword
		}
		fNewPassword := frNewPassword
		if fNewPassword != "" {
			if err := r.SetFormParam("newPassword", fNewPassword); err != nil {
				return err
			}
		}

	}

	if o.NewPasswordConfirm != nil {

		// form param newPasswordConfirm
		var frNewPasswordConfirm string
		if o.NewPasswordConfirm != nil {
			frNewPasswordConfirm = *o.NewPasswordConfirm
		}
		fNewPasswordConfirm := frNewPasswordConfirm
		if fNewPasswordConfirm != "" {
			if err := r.SetFormParam("newPasswordConfirm", fNewPasswordConfirm); err != nil {
				return err
			}
		}

	}

	if o.OldPassword != nil {

		// form param oldPassword
		var frOldPassword string
		if o.OldPassword != nil {
			frOldPassword = *o.OldPassword
		}
		fOldPassword := frOldPassword
		if fOldPassword != "" {
			if err := r.SetFormParam("oldPassword", fOldPassword); err != nil {
				return err
			}
		}

	}

	if o.PgpPubKey != nil {

		// form param pgpPubKey
		var frPgpPubKey string
		if o.PgpPubKey != nil {
			frPgpPubKey = *o.PgpPubKey
		}
		fPgpPubKey := frPgpPubKey
		if fPgpPubKey != "" {
			if err := r.SetFormParam("pgpPubKey", fPgpPubKey); err != nil {
				return err
			}
		}

	}

	if o.Username != nil {

		// form param username
		var frUsername string
		if o.Username != nil {
			frUsername = *o.Username
		}
		fUsername := frUsername
		if fUsername != "" {
			if err := r.SetFormParam("username", fUsername); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
