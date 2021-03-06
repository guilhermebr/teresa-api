package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDeploymentsParams creates a new GetDeploymentsParams object
// with the default values initialized.
func NewGetDeploymentsParams() *GetDeploymentsParams {
	var (
		limitDefault = int64(20)
		sinceDefault = int64(0)
	)
	return &GetDeploymentsParams{
		Limit: &limitDefault,
		Since: &sinceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentsParamsWithTimeout creates a new GetDeploymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentsParamsWithTimeout(timeout time.Duration) *GetDeploymentsParams {
	var (
		limitDefault int64 = int64(20)
		sinceDefault int64 = int64(0)
	)
	return &GetDeploymentsParams{
		Limit: &limitDefault,
		Since: &sinceDefault,

		timeout: timeout,
	}
}

/*GetDeploymentsParams contains all the parameters to send to the API endpoint
for the get deployments operation typically these are written to a http.Request
*/
type GetDeploymentsParams struct {

	/*AppID
	  App ID

	*/
	AppID int64
	/*Limit
	  Limit

	*/
	Limit *int64
	/*Since
	  Offset

	*/
	Since *int64
	/*TeamID
	  Team ID

	*/
	TeamID int64

	timeout time.Duration
}

// WithAppID adds the appID to the get deployments params
func (o *GetDeploymentsParams) WithAppID(appID int64) *GetDeploymentsParams {
	o.AppID = appID
	return o
}

// WithLimit adds the limit to the get deployments params
func (o *GetDeploymentsParams) WithLimit(limit *int64) *GetDeploymentsParams {
	o.Limit = limit
	return o
}

// WithSince adds the since to the get deployments params
func (o *GetDeploymentsParams) WithSince(since *int64) *GetDeploymentsParams {
	o.Since = since
	return o
}

// WithTeamID adds the teamID to the get deployments params
func (o *GetDeploymentsParams) WithTeamID(teamID int64) *GetDeploymentsParams {
	o.TeamID = teamID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", swag.FormatInt64(o.AppID)); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Since != nil {

		// query param since
		var qrSince int64
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := swag.FormatInt64(qrSince)
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	// path param team_id
	if err := r.SetPathParam("team_id", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
