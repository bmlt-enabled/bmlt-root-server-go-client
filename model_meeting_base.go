/*
BMLT

BMLT Admin API Documentation

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmlt

import (
	"encoding/json"
)

// MeetingBase struct for MeetingBase
type MeetingBase struct {
	ServiceBodyId *int32 `json:"serviceBodyId,omitempty"`
	FormatIds []int32 `json:"formatIds,omitempty"`
	VenueType *int32 `json:"venueType,omitempty"`
	TemporarilyVirtual *bool `json:"temporarilyVirtual,omitempty"`
	Day *int32 `json:"day,omitempty"`
	StartTime *string `json:"startTime,omitempty"`
	Duration *string `json:"duration,omitempty"`
	TimeZone *string `json:"timeZone,omitempty"`
	Latitude *float32 `json:"latitude,omitempty"`
	Longitude *float32 `json:"longitude,omitempty"`
	Published *bool `json:"published,omitempty"`
	Email *string `json:"email,omitempty"`
	WorldId *string `json:"worldId,omitempty"`
	Name *string `json:"name,omitempty"`
	LocationText *string `json:"location_text,omitempty"`
	LocationInfo *string `json:"location_info,omitempty"`
	LocationStreet *string `json:"location_street,omitempty"`
	LocationNeighborhood *string `json:"location_neighborhood,omitempty"`
	LocationCitySubsection *string `json:"location_city_subsection,omitempty"`
	LocationMunicipality *string `json:"location_municipality,omitempty"`
	LocationSubProvince *string `json:"location_sub_province,omitempty"`
	LocationProvince *string `json:"location_province,omitempty"`
	LocationPostalCode1 *string `json:"location_postal_code_1,omitempty"`
	LocationNation *string `json:"location_nation,omitempty"`
	PhoneMeetingNumber *string `json:"phone_meeting_number,omitempty"`
	VirtualMeetingLink *string `json:"virtual_meeting_link,omitempty"`
	VirtualMeetingAdditionalInfo *string `json:"virtual_meeting_additional_info,omitempty"`
	ContactName1 *string `json:"contact_name_1,omitempty"`
	ContactName2 *string `json:"contact_name_2,omitempty"`
	ContactPhone1 *string `json:"contact_phone_1,omitempty"`
	ContactPhone2 *string `json:"contact_phone_2,omitempty"`
	ContactEmail1 *string `json:"contact_email_1,omitempty"`
	ContactEmail2 *string `json:"contact_email_2,omitempty"`
	BusLines *string `json:"bus_lines,omitempty"`
	TrainLine *string `json:"train_line,omitempty"`
	Comments *string `json:"comments,omitempty"`
}

// NewMeetingBase instantiates a new MeetingBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeetingBase() *MeetingBase {
	this := MeetingBase{}
	return &this
}

// NewMeetingBaseWithDefaults instantiates a new MeetingBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeetingBaseWithDefaults() *MeetingBase {
	this := MeetingBase{}
	return &this
}

// GetServiceBodyId returns the ServiceBodyId field value if set, zero value otherwise.
func (o *MeetingBase) GetServiceBodyId() int32 {
	if o == nil || o.ServiceBodyId == nil {
		var ret int32
		return ret
	}
	return *o.ServiceBodyId
}

// GetServiceBodyIdOk returns a tuple with the ServiceBodyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetServiceBodyIdOk() (*int32, bool) {
	if o == nil || o.ServiceBodyId == nil {
		return nil, false
	}
	return o.ServiceBodyId, true
}

// HasServiceBodyId returns a boolean if a field has been set.
func (o *MeetingBase) HasServiceBodyId() bool {
	if o != nil && o.ServiceBodyId != nil {
		return true
	}

	return false
}

// SetServiceBodyId gets a reference to the given int32 and assigns it to the ServiceBodyId field.
func (o *MeetingBase) SetServiceBodyId(v int32) {
	o.ServiceBodyId = &v
}

// GetFormatIds returns the FormatIds field value if set, zero value otherwise.
func (o *MeetingBase) GetFormatIds() []int32 {
	if o == nil || o.FormatIds == nil {
		var ret []int32
		return ret
	}
	return o.FormatIds
}

// GetFormatIdsOk returns a tuple with the FormatIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetFormatIdsOk() ([]int32, bool) {
	if o == nil || o.FormatIds == nil {
		return nil, false
	}
	return o.FormatIds, true
}

// HasFormatIds returns a boolean if a field has been set.
func (o *MeetingBase) HasFormatIds() bool {
	if o != nil && o.FormatIds != nil {
		return true
	}

	return false
}

// SetFormatIds gets a reference to the given []int32 and assigns it to the FormatIds field.
func (o *MeetingBase) SetFormatIds(v []int32) {
	o.FormatIds = v
}

// GetVenueType returns the VenueType field value if set, zero value otherwise.
func (o *MeetingBase) GetVenueType() int32 {
	if o == nil || o.VenueType == nil {
		var ret int32
		return ret
	}
	return *o.VenueType
}

// GetVenueTypeOk returns a tuple with the VenueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetVenueTypeOk() (*int32, bool) {
	if o == nil || o.VenueType == nil {
		return nil, false
	}
	return o.VenueType, true
}

// HasVenueType returns a boolean if a field has been set.
func (o *MeetingBase) HasVenueType() bool {
	if o != nil && o.VenueType != nil {
		return true
	}

	return false
}

// SetVenueType gets a reference to the given int32 and assigns it to the VenueType field.
func (o *MeetingBase) SetVenueType(v int32) {
	o.VenueType = &v
}

// GetTemporarilyVirtual returns the TemporarilyVirtual field value if set, zero value otherwise.
func (o *MeetingBase) GetTemporarilyVirtual() bool {
	if o == nil || o.TemporarilyVirtual == nil {
		var ret bool
		return ret
	}
	return *o.TemporarilyVirtual
}

// GetTemporarilyVirtualOk returns a tuple with the TemporarilyVirtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetTemporarilyVirtualOk() (*bool, bool) {
	if o == nil || o.TemporarilyVirtual == nil {
		return nil, false
	}
	return o.TemporarilyVirtual, true
}

// HasTemporarilyVirtual returns a boolean if a field has been set.
func (o *MeetingBase) HasTemporarilyVirtual() bool {
	if o != nil && o.TemporarilyVirtual != nil {
		return true
	}

	return false
}

// SetTemporarilyVirtual gets a reference to the given bool and assigns it to the TemporarilyVirtual field.
func (o *MeetingBase) SetTemporarilyVirtual(v bool) {
	o.TemporarilyVirtual = &v
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *MeetingBase) GetDay() int32 {
	if o == nil || o.Day == nil {
		var ret int32
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetDayOk() (*int32, bool) {
	if o == nil || o.Day == nil {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *MeetingBase) HasDay() bool {
	if o != nil && o.Day != nil {
		return true
	}

	return false
}

// SetDay gets a reference to the given int32 and assigns it to the Day field.
func (o *MeetingBase) SetDay(v int32) {
	o.Day = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *MeetingBase) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *MeetingBase) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *MeetingBase) SetStartTime(v string) {
	o.StartTime = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *MeetingBase) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *MeetingBase) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *MeetingBase) SetDuration(v string) {
	o.Duration = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *MeetingBase) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *MeetingBase) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *MeetingBase) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *MeetingBase) GetLatitude() float32 {
	if o == nil || o.Latitude == nil {
		var ret float32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetLatitudeOk() (*float32, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *MeetingBase) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float32 and assigns it to the Latitude field.
func (o *MeetingBase) SetLatitude(v float32) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *MeetingBase) GetLongitude() float32 {
	if o == nil || o.Longitude == nil {
		var ret float32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetLongitudeOk() (*float32, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *MeetingBase) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float32 and assigns it to the Longitude field.
func (o *MeetingBase) SetLongitude(v float32) {
	o.Longitude = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *MeetingBase) GetPublished() bool {
	if o == nil || o.Published == nil {
		var ret bool
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetPublishedOk() (*bool, bool) {
	if o == nil || o.Published == nil {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *MeetingBase) HasPublished() bool {
	if o != nil && o.Published != nil {
		return true
	}

	return false
}

// SetPublished gets a reference to the given bool and assigns it to the Published field.
func (o *MeetingBase) SetPublished(v bool) {
	o.Published = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *MeetingBase) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MeetingBase) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MeetingBase) SetEmail(v string) {
	o.Email = &v
}

// GetWorldId returns the WorldId field value if set, zero value otherwise.
func (o *MeetingBase) GetWorldId() string {
	if o == nil || o.WorldId == nil {
		var ret string
		return ret
	}
	return *o.WorldId
}

// GetWorldIdOk returns a tuple with the WorldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetWorldIdOk() (*string, bool) {
	if o == nil || o.WorldId == nil {
		return nil, false
	}
	return o.WorldId, true
}

// HasWorldId returns a boolean if a field has been set.
func (o *MeetingBase) HasWorldId() bool {
	if o != nil && o.WorldId != nil {
		return true
	}

	return false
}

// SetWorldId gets a reference to the given string and assigns it to the WorldId field.
func (o *MeetingBase) SetWorldId(v string) {
	o.WorldId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MeetingBase) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MeetingBase) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MeetingBase) SetName(v string) {
	o.Name = &v
}

// GetLocationText returns the LocationText field value if set, zero value otherwise.
func (o *MeetingBase) GetLocationText() string {
	if o == nil || o.LocationText == nil {
		var ret string
		return ret
	}
	return *o.LocationText
}

// GetLocationTextOk returns a tuple with the LocationText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetLocationTextOk() (*string, bool) {
	if o == nil || o.LocationText == nil {
		return nil, false
	}
	return o.LocationText, true
}

// HasLocationText returns a boolean if a field has been set.
func (o *MeetingBase) HasLocationText() bool {
	if o != nil && o.LocationText != nil {
		return true
	}

	return false
}

// SetLocationText gets a reference to the given string and assigns it to the LocationText field.
func (o *MeetingBase) SetLocationText(v string) {
	o.LocationText = &v
}

// GetLocationInfo returns the LocationInfo field value if set, zero value otherwise.
func (o *MeetingBase) GetLocationInfo() string {
	if o == nil || o.LocationInfo == nil {
		var ret string
		return ret
	}
	return *o.LocationInfo
}

// GetLocationInfoOk returns a tuple with the LocationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetLocationInfoOk() (*string, bool) {
	if o == nil || o.LocationInfo == nil {
		return nil, false
	}
	return o.LocationInfo, true
}

// HasLocationInfo returns a boolean if a field has been set.
func (o *MeetingBase) HasLocationInfo() bool {
	if o != nil && o.LocationInfo != nil {
		return true
	}

	return false
}

// SetLocationInfo gets a reference to the given string and assigns it to the LocationInfo field.
func (o *MeetingBase) SetLocationInfo(v string) {
	o.LocationInfo = &v
}

// GetLocationStreet returns the LocationStreet field value if set, zero value otherwise.
func (o *MeetingBase) GetLocationStreet() string {
	if o == nil || o.LocationStreet == nil {
		var ret string
		return ret
	}
	return *o.LocationStreet
}

// GetLocationStreetOk returns a tuple with the LocationStreet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetLocationStreetOk() (*string, bool) {
	if o == nil || o.LocationStreet == nil {
		return nil, false
	}
	return o.LocationStreet, true
}

// HasLocationStreet returns a boolean if a field has been set.
func (o *MeetingBase) HasLocationStreet() bool {
	if o != nil && o.LocationStreet != nil {
		return true
	}

	return false
}

// SetLocationStreet gets a reference to the given string and assigns it to the LocationStreet field.
func (o *MeetingBase) SetLocationStreet(v string) {
	o.LocationStreet = &v
}

// GetLocationNeighborhood returns the LocationNeighborhood field value if set, zero value otherwise.
func (o *MeetingBase) GetLocationNeighborhood() string {
	if o == nil || o.LocationNeighborhood == nil {
		var ret string
		return ret
	}
	return *o.LocationNeighborhood
}

// GetLocationNeighborhoodOk returns a tuple with the LocationNeighborhood field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetLocationNeighborhoodOk() (*string, bool) {
	if o == nil || o.LocationNeighborhood == nil {
		return nil, false
	}
	return o.LocationNeighborhood, true
}

// HasLocationNeighborhood returns a boolean if a field has been set.
func (o *MeetingBase) HasLocationNeighborhood() bool {
	if o != nil && o.LocationNeighborhood != nil {
		return true
	}

	return false
}

// SetLocationNeighborhood gets a reference to the given string and assigns it to the LocationNeighborhood field.
func (o *MeetingBase) SetLocationNeighborhood(v string) {
	o.LocationNeighborhood = &v
}

// GetLocationCitySubsection returns the LocationCitySubsection field value if set, zero value otherwise.
func (o *MeetingBase) GetLocationCitySubsection() string {
	if o == nil || o.LocationCitySubsection == nil {
		var ret string
		return ret
	}
	return *o.LocationCitySubsection
}

// GetLocationCitySubsectionOk returns a tuple with the LocationCitySubsection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetLocationCitySubsectionOk() (*string, bool) {
	if o == nil || o.LocationCitySubsection == nil {
		return nil, false
	}
	return o.LocationCitySubsection, true
}

// HasLocationCitySubsection returns a boolean if a field has been set.
func (o *MeetingBase) HasLocationCitySubsection() bool {
	if o != nil && o.LocationCitySubsection != nil {
		return true
	}

	return false
}

// SetLocationCitySubsection gets a reference to the given string and assigns it to the LocationCitySubsection field.
func (o *MeetingBase) SetLocationCitySubsection(v string) {
	o.LocationCitySubsection = &v
}

// GetLocationMunicipality returns the LocationMunicipality field value if set, zero value otherwise.
func (o *MeetingBase) GetLocationMunicipality() string {
	if o == nil || o.LocationMunicipality == nil {
		var ret string
		return ret
	}
	return *o.LocationMunicipality
}

// GetLocationMunicipalityOk returns a tuple with the LocationMunicipality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetLocationMunicipalityOk() (*string, bool) {
	if o == nil || o.LocationMunicipality == nil {
		return nil, false
	}
	return o.LocationMunicipality, true
}

// HasLocationMunicipality returns a boolean if a field has been set.
func (o *MeetingBase) HasLocationMunicipality() bool {
	if o != nil && o.LocationMunicipality != nil {
		return true
	}

	return false
}

// SetLocationMunicipality gets a reference to the given string and assigns it to the LocationMunicipality field.
func (o *MeetingBase) SetLocationMunicipality(v string) {
	o.LocationMunicipality = &v
}

// GetLocationSubProvince returns the LocationSubProvince field value if set, zero value otherwise.
func (o *MeetingBase) GetLocationSubProvince() string {
	if o == nil || o.LocationSubProvince == nil {
		var ret string
		return ret
	}
	return *o.LocationSubProvince
}

// GetLocationSubProvinceOk returns a tuple with the LocationSubProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetLocationSubProvinceOk() (*string, bool) {
	if o == nil || o.LocationSubProvince == nil {
		return nil, false
	}
	return o.LocationSubProvince, true
}

// HasLocationSubProvince returns a boolean if a field has been set.
func (o *MeetingBase) HasLocationSubProvince() bool {
	if o != nil && o.LocationSubProvince != nil {
		return true
	}

	return false
}

// SetLocationSubProvince gets a reference to the given string and assigns it to the LocationSubProvince field.
func (o *MeetingBase) SetLocationSubProvince(v string) {
	o.LocationSubProvince = &v
}

// GetLocationProvince returns the LocationProvince field value if set, zero value otherwise.
func (o *MeetingBase) GetLocationProvince() string {
	if o == nil || o.LocationProvince == nil {
		var ret string
		return ret
	}
	return *o.LocationProvince
}

// GetLocationProvinceOk returns a tuple with the LocationProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetLocationProvinceOk() (*string, bool) {
	if o == nil || o.LocationProvince == nil {
		return nil, false
	}
	return o.LocationProvince, true
}

// HasLocationProvince returns a boolean if a field has been set.
func (o *MeetingBase) HasLocationProvince() bool {
	if o != nil && o.LocationProvince != nil {
		return true
	}

	return false
}

// SetLocationProvince gets a reference to the given string and assigns it to the LocationProvince field.
func (o *MeetingBase) SetLocationProvince(v string) {
	o.LocationProvince = &v
}

// GetLocationPostalCode1 returns the LocationPostalCode1 field value if set, zero value otherwise.
func (o *MeetingBase) GetLocationPostalCode1() string {
	if o == nil || o.LocationPostalCode1 == nil {
		var ret string
		return ret
	}
	return *o.LocationPostalCode1
}

// GetLocationPostalCode1Ok returns a tuple with the LocationPostalCode1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetLocationPostalCode1Ok() (*string, bool) {
	if o == nil || o.LocationPostalCode1 == nil {
		return nil, false
	}
	return o.LocationPostalCode1, true
}

// HasLocationPostalCode1 returns a boolean if a field has been set.
func (o *MeetingBase) HasLocationPostalCode1() bool {
	if o != nil && o.LocationPostalCode1 != nil {
		return true
	}

	return false
}

// SetLocationPostalCode1 gets a reference to the given string and assigns it to the LocationPostalCode1 field.
func (o *MeetingBase) SetLocationPostalCode1(v string) {
	o.LocationPostalCode1 = &v
}

// GetLocationNation returns the LocationNation field value if set, zero value otherwise.
func (o *MeetingBase) GetLocationNation() string {
	if o == nil || o.LocationNation == nil {
		var ret string
		return ret
	}
	return *o.LocationNation
}

// GetLocationNationOk returns a tuple with the LocationNation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetLocationNationOk() (*string, bool) {
	if o == nil || o.LocationNation == nil {
		return nil, false
	}
	return o.LocationNation, true
}

// HasLocationNation returns a boolean if a field has been set.
func (o *MeetingBase) HasLocationNation() bool {
	if o != nil && o.LocationNation != nil {
		return true
	}

	return false
}

// SetLocationNation gets a reference to the given string and assigns it to the LocationNation field.
func (o *MeetingBase) SetLocationNation(v string) {
	o.LocationNation = &v
}

// GetPhoneMeetingNumber returns the PhoneMeetingNumber field value if set, zero value otherwise.
func (o *MeetingBase) GetPhoneMeetingNumber() string {
	if o == nil || o.PhoneMeetingNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneMeetingNumber
}

// GetPhoneMeetingNumberOk returns a tuple with the PhoneMeetingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetPhoneMeetingNumberOk() (*string, bool) {
	if o == nil || o.PhoneMeetingNumber == nil {
		return nil, false
	}
	return o.PhoneMeetingNumber, true
}

// HasPhoneMeetingNumber returns a boolean if a field has been set.
func (o *MeetingBase) HasPhoneMeetingNumber() bool {
	if o != nil && o.PhoneMeetingNumber != nil {
		return true
	}

	return false
}

// SetPhoneMeetingNumber gets a reference to the given string and assigns it to the PhoneMeetingNumber field.
func (o *MeetingBase) SetPhoneMeetingNumber(v string) {
	o.PhoneMeetingNumber = &v
}

// GetVirtualMeetingLink returns the VirtualMeetingLink field value if set, zero value otherwise.
func (o *MeetingBase) GetVirtualMeetingLink() string {
	if o == nil || o.VirtualMeetingLink == nil {
		var ret string
		return ret
	}
	return *o.VirtualMeetingLink
}

// GetVirtualMeetingLinkOk returns a tuple with the VirtualMeetingLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetVirtualMeetingLinkOk() (*string, bool) {
	if o == nil || o.VirtualMeetingLink == nil {
		return nil, false
	}
	return o.VirtualMeetingLink, true
}

// HasVirtualMeetingLink returns a boolean if a field has been set.
func (o *MeetingBase) HasVirtualMeetingLink() bool {
	if o != nil && o.VirtualMeetingLink != nil {
		return true
	}

	return false
}

// SetVirtualMeetingLink gets a reference to the given string and assigns it to the VirtualMeetingLink field.
func (o *MeetingBase) SetVirtualMeetingLink(v string) {
	o.VirtualMeetingLink = &v
}

// GetVirtualMeetingAdditionalInfo returns the VirtualMeetingAdditionalInfo field value if set, zero value otherwise.
func (o *MeetingBase) GetVirtualMeetingAdditionalInfo() string {
	if o == nil || o.VirtualMeetingAdditionalInfo == nil {
		var ret string
		return ret
	}
	return *o.VirtualMeetingAdditionalInfo
}

// GetVirtualMeetingAdditionalInfoOk returns a tuple with the VirtualMeetingAdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetVirtualMeetingAdditionalInfoOk() (*string, bool) {
	if o == nil || o.VirtualMeetingAdditionalInfo == nil {
		return nil, false
	}
	return o.VirtualMeetingAdditionalInfo, true
}

// HasVirtualMeetingAdditionalInfo returns a boolean if a field has been set.
func (o *MeetingBase) HasVirtualMeetingAdditionalInfo() bool {
	if o != nil && o.VirtualMeetingAdditionalInfo != nil {
		return true
	}

	return false
}

// SetVirtualMeetingAdditionalInfo gets a reference to the given string and assigns it to the VirtualMeetingAdditionalInfo field.
func (o *MeetingBase) SetVirtualMeetingAdditionalInfo(v string) {
	o.VirtualMeetingAdditionalInfo = &v
}

// GetContactName1 returns the ContactName1 field value if set, zero value otherwise.
func (o *MeetingBase) GetContactName1() string {
	if o == nil || o.ContactName1 == nil {
		var ret string
		return ret
	}
	return *o.ContactName1
}

// GetContactName1Ok returns a tuple with the ContactName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetContactName1Ok() (*string, bool) {
	if o == nil || o.ContactName1 == nil {
		return nil, false
	}
	return o.ContactName1, true
}

// HasContactName1 returns a boolean if a field has been set.
func (o *MeetingBase) HasContactName1() bool {
	if o != nil && o.ContactName1 != nil {
		return true
	}

	return false
}

// SetContactName1 gets a reference to the given string and assigns it to the ContactName1 field.
func (o *MeetingBase) SetContactName1(v string) {
	o.ContactName1 = &v
}

// GetContactName2 returns the ContactName2 field value if set, zero value otherwise.
func (o *MeetingBase) GetContactName2() string {
	if o == nil || o.ContactName2 == nil {
		var ret string
		return ret
	}
	return *o.ContactName2
}

// GetContactName2Ok returns a tuple with the ContactName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetContactName2Ok() (*string, bool) {
	if o == nil || o.ContactName2 == nil {
		return nil, false
	}
	return o.ContactName2, true
}

// HasContactName2 returns a boolean if a field has been set.
func (o *MeetingBase) HasContactName2() bool {
	if o != nil && o.ContactName2 != nil {
		return true
	}

	return false
}

// SetContactName2 gets a reference to the given string and assigns it to the ContactName2 field.
func (o *MeetingBase) SetContactName2(v string) {
	o.ContactName2 = &v
}

// GetContactPhone1 returns the ContactPhone1 field value if set, zero value otherwise.
func (o *MeetingBase) GetContactPhone1() string {
	if o == nil || o.ContactPhone1 == nil {
		var ret string
		return ret
	}
	return *o.ContactPhone1
}

// GetContactPhone1Ok returns a tuple with the ContactPhone1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetContactPhone1Ok() (*string, bool) {
	if o == nil || o.ContactPhone1 == nil {
		return nil, false
	}
	return o.ContactPhone1, true
}

// HasContactPhone1 returns a boolean if a field has been set.
func (o *MeetingBase) HasContactPhone1() bool {
	if o != nil && o.ContactPhone1 != nil {
		return true
	}

	return false
}

// SetContactPhone1 gets a reference to the given string and assigns it to the ContactPhone1 field.
func (o *MeetingBase) SetContactPhone1(v string) {
	o.ContactPhone1 = &v
}

// GetContactPhone2 returns the ContactPhone2 field value if set, zero value otherwise.
func (o *MeetingBase) GetContactPhone2() string {
	if o == nil || o.ContactPhone2 == nil {
		var ret string
		return ret
	}
	return *o.ContactPhone2
}

// GetContactPhone2Ok returns a tuple with the ContactPhone2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetContactPhone2Ok() (*string, bool) {
	if o == nil || o.ContactPhone2 == nil {
		return nil, false
	}
	return o.ContactPhone2, true
}

// HasContactPhone2 returns a boolean if a field has been set.
func (o *MeetingBase) HasContactPhone2() bool {
	if o != nil && o.ContactPhone2 != nil {
		return true
	}

	return false
}

// SetContactPhone2 gets a reference to the given string and assigns it to the ContactPhone2 field.
func (o *MeetingBase) SetContactPhone2(v string) {
	o.ContactPhone2 = &v
}

// GetContactEmail1 returns the ContactEmail1 field value if set, zero value otherwise.
func (o *MeetingBase) GetContactEmail1() string {
	if o == nil || o.ContactEmail1 == nil {
		var ret string
		return ret
	}
	return *o.ContactEmail1
}

// GetContactEmail1Ok returns a tuple with the ContactEmail1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetContactEmail1Ok() (*string, bool) {
	if o == nil || o.ContactEmail1 == nil {
		return nil, false
	}
	return o.ContactEmail1, true
}

// HasContactEmail1 returns a boolean if a field has been set.
func (o *MeetingBase) HasContactEmail1() bool {
	if o != nil && o.ContactEmail1 != nil {
		return true
	}

	return false
}

// SetContactEmail1 gets a reference to the given string and assigns it to the ContactEmail1 field.
func (o *MeetingBase) SetContactEmail1(v string) {
	o.ContactEmail1 = &v
}

// GetContactEmail2 returns the ContactEmail2 field value if set, zero value otherwise.
func (o *MeetingBase) GetContactEmail2() string {
	if o == nil || o.ContactEmail2 == nil {
		var ret string
		return ret
	}
	return *o.ContactEmail2
}

// GetContactEmail2Ok returns a tuple with the ContactEmail2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetContactEmail2Ok() (*string, bool) {
	if o == nil || o.ContactEmail2 == nil {
		return nil, false
	}
	return o.ContactEmail2, true
}

// HasContactEmail2 returns a boolean if a field has been set.
func (o *MeetingBase) HasContactEmail2() bool {
	if o != nil && o.ContactEmail2 != nil {
		return true
	}

	return false
}

// SetContactEmail2 gets a reference to the given string and assigns it to the ContactEmail2 field.
func (o *MeetingBase) SetContactEmail2(v string) {
	o.ContactEmail2 = &v
}

// GetBusLines returns the BusLines field value if set, zero value otherwise.
func (o *MeetingBase) GetBusLines() string {
	if o == nil || o.BusLines == nil {
		var ret string
		return ret
	}
	return *o.BusLines
}

// GetBusLinesOk returns a tuple with the BusLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetBusLinesOk() (*string, bool) {
	if o == nil || o.BusLines == nil {
		return nil, false
	}
	return o.BusLines, true
}

// HasBusLines returns a boolean if a field has been set.
func (o *MeetingBase) HasBusLines() bool {
	if o != nil && o.BusLines != nil {
		return true
	}

	return false
}

// SetBusLines gets a reference to the given string and assigns it to the BusLines field.
func (o *MeetingBase) SetBusLines(v string) {
	o.BusLines = &v
}

// GetTrainLine returns the TrainLine field value if set, zero value otherwise.
func (o *MeetingBase) GetTrainLine() string {
	if o == nil || o.TrainLine == nil {
		var ret string
		return ret
	}
	return *o.TrainLine
}

// GetTrainLineOk returns a tuple with the TrainLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetTrainLineOk() (*string, bool) {
	if o == nil || o.TrainLine == nil {
		return nil, false
	}
	return o.TrainLine, true
}

// HasTrainLine returns a boolean if a field has been set.
func (o *MeetingBase) HasTrainLine() bool {
	if o != nil && o.TrainLine != nil {
		return true
	}

	return false
}

// SetTrainLine gets a reference to the given string and assigns it to the TrainLine field.
func (o *MeetingBase) SetTrainLine(v string) {
	o.TrainLine = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *MeetingBase) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingBase) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *MeetingBase) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *MeetingBase) SetComments(v string) {
	o.Comments = &v
}

func (o MeetingBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceBodyId != nil {
		toSerialize["serviceBodyId"] = o.ServiceBodyId
	}
	if o.FormatIds != nil {
		toSerialize["formatIds"] = o.FormatIds
	}
	if o.VenueType != nil {
		toSerialize["venueType"] = o.VenueType
	}
	if o.TemporarilyVirtual != nil {
		toSerialize["temporarilyVirtual"] = o.TemporarilyVirtual
	}
	if o.Day != nil {
		toSerialize["day"] = o.Day
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}
	if o.Published != nil {
		toSerialize["published"] = o.Published
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.WorldId != nil {
		toSerialize["worldId"] = o.WorldId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.LocationText != nil {
		toSerialize["location_text"] = o.LocationText
	}
	if o.LocationInfo != nil {
		toSerialize["location_info"] = o.LocationInfo
	}
	if o.LocationStreet != nil {
		toSerialize["location_street"] = o.LocationStreet
	}
	if o.LocationNeighborhood != nil {
		toSerialize["location_neighborhood"] = o.LocationNeighborhood
	}
	if o.LocationCitySubsection != nil {
		toSerialize["location_city_subsection"] = o.LocationCitySubsection
	}
	if o.LocationMunicipality != nil {
		toSerialize["location_municipality"] = o.LocationMunicipality
	}
	if o.LocationSubProvince != nil {
		toSerialize["location_sub_province"] = o.LocationSubProvince
	}
	if o.LocationProvince != nil {
		toSerialize["location_province"] = o.LocationProvince
	}
	if o.LocationPostalCode1 != nil {
		toSerialize["location_postal_code_1"] = o.LocationPostalCode1
	}
	if o.LocationNation != nil {
		toSerialize["location_nation"] = o.LocationNation
	}
	if o.PhoneMeetingNumber != nil {
		toSerialize["phone_meeting_number"] = o.PhoneMeetingNumber
	}
	if o.VirtualMeetingLink != nil {
		toSerialize["virtual_meeting_link"] = o.VirtualMeetingLink
	}
	if o.VirtualMeetingAdditionalInfo != nil {
		toSerialize["virtual_meeting_additional_info"] = o.VirtualMeetingAdditionalInfo
	}
	if o.ContactName1 != nil {
		toSerialize["contact_name_1"] = o.ContactName1
	}
	if o.ContactName2 != nil {
		toSerialize["contact_name_2"] = o.ContactName2
	}
	if o.ContactPhone1 != nil {
		toSerialize["contact_phone_1"] = o.ContactPhone1
	}
	if o.ContactPhone2 != nil {
		toSerialize["contact_phone_2"] = o.ContactPhone2
	}
	if o.ContactEmail1 != nil {
		toSerialize["contact_email_1"] = o.ContactEmail1
	}
	if o.ContactEmail2 != nil {
		toSerialize["contact_email_2"] = o.ContactEmail2
	}
	if o.BusLines != nil {
		toSerialize["bus_lines"] = o.BusLines
	}
	if o.TrainLine != nil {
		toSerialize["train_line"] = o.TrainLine
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	return json.Marshal(toSerialize)
}

type NullableMeetingBase struct {
	value *MeetingBase
	isSet bool
}

func (v NullableMeetingBase) Get() *MeetingBase {
	return v.value
}

func (v *NullableMeetingBase) Set(val *MeetingBase) {
	v.value = val
	v.isSet = true
}

func (v NullableMeetingBase) IsSet() bool {
	return v.isSet
}

func (v *NullableMeetingBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeetingBase(val *MeetingBase) *NullableMeetingBase {
	return &NullableMeetingBase{value: val, isSet: true}
}

func (v NullableMeetingBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeetingBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


