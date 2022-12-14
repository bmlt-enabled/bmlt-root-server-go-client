# MeetingBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceBodyId** | Pointer to **int32** |  | [optional] 
**FormatIds** | Pointer to **[]int32** |  | [optional] 
**VenueType** | Pointer to **int32** |  | [optional] 
**TemporarilyVirtual** | Pointer to **bool** |  | [optional] 
**Day** | Pointer to **int32** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**Latitude** | Pointer to **float32** |  | [optional] 
**Longitude** | Pointer to **float32** |  | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**WorldId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LocationText** | Pointer to **string** |  | [optional] 
**LocationInfo** | Pointer to **string** |  | [optional] 
**LocationStreet** | Pointer to **string** |  | [optional] 
**LocationNeighborhood** | Pointer to **string** |  | [optional] 
**LocationCitySubsection** | Pointer to **string** |  | [optional] 
**LocationMunicipality** | Pointer to **string** |  | [optional] 
**LocationSubProvince** | Pointer to **string** |  | [optional] 
**LocationProvince** | Pointer to **string** |  | [optional] 
**LocationPostalCode1** | Pointer to **string** |  | [optional] 
**LocationNation** | Pointer to **string** |  | [optional] 
**PhoneMeetingNumber** | Pointer to **string** |  | [optional] 
**VirtualMeetingLink** | Pointer to **string** |  | [optional] 
**VirtualMeetingAdditionalInfo** | Pointer to **string** |  | [optional] 
**ContactName1** | Pointer to **string** |  | [optional] 
**ContactName2** | Pointer to **string** |  | [optional] 
**ContactPhone1** | Pointer to **string** |  | [optional] 
**ContactPhone2** | Pointer to **string** |  | [optional] 
**ContactEmail1** | Pointer to **string** |  | [optional] 
**ContactEmail2** | Pointer to **string** |  | [optional] 
**BusLines** | Pointer to **string** |  | [optional] 
**TrainLine** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 

## Methods

### NewMeetingBase

`func NewMeetingBase() *MeetingBase`

NewMeetingBase instantiates a new MeetingBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingBaseWithDefaults

`func NewMeetingBaseWithDefaults() *MeetingBase`

NewMeetingBaseWithDefaults instantiates a new MeetingBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceBodyId

`func (o *MeetingBase) GetServiceBodyId() int32`

GetServiceBodyId returns the ServiceBodyId field if non-nil, zero value otherwise.

### GetServiceBodyIdOk

`func (o *MeetingBase) GetServiceBodyIdOk() (*int32, bool)`

GetServiceBodyIdOk returns a tuple with the ServiceBodyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBodyId

`func (o *MeetingBase) SetServiceBodyId(v int32)`

SetServiceBodyId sets ServiceBodyId field to given value.

### HasServiceBodyId

`func (o *MeetingBase) HasServiceBodyId() bool`

HasServiceBodyId returns a boolean if a field has been set.

### GetFormatIds

`func (o *MeetingBase) GetFormatIds() []int32`

GetFormatIds returns the FormatIds field if non-nil, zero value otherwise.

### GetFormatIdsOk

`func (o *MeetingBase) GetFormatIdsOk() (*[]int32, bool)`

GetFormatIdsOk returns a tuple with the FormatIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatIds

`func (o *MeetingBase) SetFormatIds(v []int32)`

SetFormatIds sets FormatIds field to given value.

### HasFormatIds

`func (o *MeetingBase) HasFormatIds() bool`

HasFormatIds returns a boolean if a field has been set.

### GetVenueType

`func (o *MeetingBase) GetVenueType() int32`

GetVenueType returns the VenueType field if non-nil, zero value otherwise.

### GetVenueTypeOk

`func (o *MeetingBase) GetVenueTypeOk() (*int32, bool)`

GetVenueTypeOk returns a tuple with the VenueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenueType

`func (o *MeetingBase) SetVenueType(v int32)`

SetVenueType sets VenueType field to given value.

### HasVenueType

`func (o *MeetingBase) HasVenueType() bool`

HasVenueType returns a boolean if a field has been set.

### GetTemporarilyVirtual

`func (o *MeetingBase) GetTemporarilyVirtual() bool`

GetTemporarilyVirtual returns the TemporarilyVirtual field if non-nil, zero value otherwise.

### GetTemporarilyVirtualOk

`func (o *MeetingBase) GetTemporarilyVirtualOk() (*bool, bool)`

GetTemporarilyVirtualOk returns a tuple with the TemporarilyVirtual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporarilyVirtual

`func (o *MeetingBase) SetTemporarilyVirtual(v bool)`

SetTemporarilyVirtual sets TemporarilyVirtual field to given value.

### HasTemporarilyVirtual

`func (o *MeetingBase) HasTemporarilyVirtual() bool`

HasTemporarilyVirtual returns a boolean if a field has been set.

### GetDay

`func (o *MeetingBase) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *MeetingBase) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *MeetingBase) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDay

`func (o *MeetingBase) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetStartTime

`func (o *MeetingBase) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MeetingBase) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MeetingBase) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MeetingBase) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetDuration

`func (o *MeetingBase) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MeetingBase) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MeetingBase) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MeetingBase) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTimeZone

`func (o *MeetingBase) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MeetingBase) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MeetingBase) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MeetingBase) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetLatitude

`func (o *MeetingBase) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *MeetingBase) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *MeetingBase) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *MeetingBase) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *MeetingBase) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *MeetingBase) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *MeetingBase) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *MeetingBase) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetPublished

`func (o *MeetingBase) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *MeetingBase) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *MeetingBase) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *MeetingBase) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetEmail

`func (o *MeetingBase) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MeetingBase) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MeetingBase) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MeetingBase) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetWorldId

`func (o *MeetingBase) GetWorldId() string`

GetWorldId returns the WorldId field if non-nil, zero value otherwise.

### GetWorldIdOk

`func (o *MeetingBase) GetWorldIdOk() (*string, bool)`

GetWorldIdOk returns a tuple with the WorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldId

`func (o *MeetingBase) SetWorldId(v string)`

SetWorldId sets WorldId field to given value.

### HasWorldId

`func (o *MeetingBase) HasWorldId() bool`

HasWorldId returns a boolean if a field has been set.

### GetName

`func (o *MeetingBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeetingBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeetingBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeetingBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocationText

`func (o *MeetingBase) GetLocationText() string`

GetLocationText returns the LocationText field if non-nil, zero value otherwise.

### GetLocationTextOk

`func (o *MeetingBase) GetLocationTextOk() (*string, bool)`

GetLocationTextOk returns a tuple with the LocationText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationText

`func (o *MeetingBase) SetLocationText(v string)`

SetLocationText sets LocationText field to given value.

### HasLocationText

`func (o *MeetingBase) HasLocationText() bool`

HasLocationText returns a boolean if a field has been set.

### GetLocationInfo

`func (o *MeetingBase) GetLocationInfo() string`

GetLocationInfo returns the LocationInfo field if non-nil, zero value otherwise.

### GetLocationInfoOk

`func (o *MeetingBase) GetLocationInfoOk() (*string, bool)`

GetLocationInfoOk returns a tuple with the LocationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationInfo

`func (o *MeetingBase) SetLocationInfo(v string)`

SetLocationInfo sets LocationInfo field to given value.

### HasLocationInfo

`func (o *MeetingBase) HasLocationInfo() bool`

HasLocationInfo returns a boolean if a field has been set.

### GetLocationStreet

`func (o *MeetingBase) GetLocationStreet() string`

GetLocationStreet returns the LocationStreet field if non-nil, zero value otherwise.

### GetLocationStreetOk

`func (o *MeetingBase) GetLocationStreetOk() (*string, bool)`

GetLocationStreetOk returns a tuple with the LocationStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationStreet

`func (o *MeetingBase) SetLocationStreet(v string)`

SetLocationStreet sets LocationStreet field to given value.

### HasLocationStreet

`func (o *MeetingBase) HasLocationStreet() bool`

HasLocationStreet returns a boolean if a field has been set.

### GetLocationNeighborhood

`func (o *MeetingBase) GetLocationNeighborhood() string`

GetLocationNeighborhood returns the LocationNeighborhood field if non-nil, zero value otherwise.

### GetLocationNeighborhoodOk

`func (o *MeetingBase) GetLocationNeighborhoodOk() (*string, bool)`

GetLocationNeighborhoodOk returns a tuple with the LocationNeighborhood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationNeighborhood

`func (o *MeetingBase) SetLocationNeighborhood(v string)`

SetLocationNeighborhood sets LocationNeighborhood field to given value.

### HasLocationNeighborhood

`func (o *MeetingBase) HasLocationNeighborhood() bool`

HasLocationNeighborhood returns a boolean if a field has been set.

### GetLocationCitySubsection

`func (o *MeetingBase) GetLocationCitySubsection() string`

GetLocationCitySubsection returns the LocationCitySubsection field if non-nil, zero value otherwise.

### GetLocationCitySubsectionOk

`func (o *MeetingBase) GetLocationCitySubsectionOk() (*string, bool)`

GetLocationCitySubsectionOk returns a tuple with the LocationCitySubsection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCitySubsection

`func (o *MeetingBase) SetLocationCitySubsection(v string)`

SetLocationCitySubsection sets LocationCitySubsection field to given value.

### HasLocationCitySubsection

`func (o *MeetingBase) HasLocationCitySubsection() bool`

HasLocationCitySubsection returns a boolean if a field has been set.

### GetLocationMunicipality

`func (o *MeetingBase) GetLocationMunicipality() string`

GetLocationMunicipality returns the LocationMunicipality field if non-nil, zero value otherwise.

### GetLocationMunicipalityOk

`func (o *MeetingBase) GetLocationMunicipalityOk() (*string, bool)`

GetLocationMunicipalityOk returns a tuple with the LocationMunicipality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationMunicipality

`func (o *MeetingBase) SetLocationMunicipality(v string)`

SetLocationMunicipality sets LocationMunicipality field to given value.

### HasLocationMunicipality

`func (o *MeetingBase) HasLocationMunicipality() bool`

HasLocationMunicipality returns a boolean if a field has been set.

### GetLocationSubProvince

`func (o *MeetingBase) GetLocationSubProvince() string`

GetLocationSubProvince returns the LocationSubProvince field if non-nil, zero value otherwise.

### GetLocationSubProvinceOk

`func (o *MeetingBase) GetLocationSubProvinceOk() (*string, bool)`

GetLocationSubProvinceOk returns a tuple with the LocationSubProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationSubProvince

`func (o *MeetingBase) SetLocationSubProvince(v string)`

SetLocationSubProvince sets LocationSubProvince field to given value.

### HasLocationSubProvince

`func (o *MeetingBase) HasLocationSubProvince() bool`

HasLocationSubProvince returns a boolean if a field has been set.

### GetLocationProvince

`func (o *MeetingBase) GetLocationProvince() string`

GetLocationProvince returns the LocationProvince field if non-nil, zero value otherwise.

### GetLocationProvinceOk

`func (o *MeetingBase) GetLocationProvinceOk() (*string, bool)`

GetLocationProvinceOk returns a tuple with the LocationProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationProvince

`func (o *MeetingBase) SetLocationProvince(v string)`

SetLocationProvince sets LocationProvince field to given value.

### HasLocationProvince

`func (o *MeetingBase) HasLocationProvince() bool`

HasLocationProvince returns a boolean if a field has been set.

### GetLocationPostalCode1

`func (o *MeetingBase) GetLocationPostalCode1() string`

GetLocationPostalCode1 returns the LocationPostalCode1 field if non-nil, zero value otherwise.

### GetLocationPostalCode1Ok

`func (o *MeetingBase) GetLocationPostalCode1Ok() (*string, bool)`

GetLocationPostalCode1Ok returns a tuple with the LocationPostalCode1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationPostalCode1

`func (o *MeetingBase) SetLocationPostalCode1(v string)`

SetLocationPostalCode1 sets LocationPostalCode1 field to given value.

### HasLocationPostalCode1

`func (o *MeetingBase) HasLocationPostalCode1() bool`

HasLocationPostalCode1 returns a boolean if a field has been set.

### GetLocationNation

`func (o *MeetingBase) GetLocationNation() string`

GetLocationNation returns the LocationNation field if non-nil, zero value otherwise.

### GetLocationNationOk

`func (o *MeetingBase) GetLocationNationOk() (*string, bool)`

GetLocationNationOk returns a tuple with the LocationNation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationNation

`func (o *MeetingBase) SetLocationNation(v string)`

SetLocationNation sets LocationNation field to given value.

### HasLocationNation

`func (o *MeetingBase) HasLocationNation() bool`

HasLocationNation returns a boolean if a field has been set.

### GetPhoneMeetingNumber

`func (o *MeetingBase) GetPhoneMeetingNumber() string`

GetPhoneMeetingNumber returns the PhoneMeetingNumber field if non-nil, zero value otherwise.

### GetPhoneMeetingNumberOk

`func (o *MeetingBase) GetPhoneMeetingNumberOk() (*string, bool)`

GetPhoneMeetingNumberOk returns a tuple with the PhoneMeetingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneMeetingNumber

`func (o *MeetingBase) SetPhoneMeetingNumber(v string)`

SetPhoneMeetingNumber sets PhoneMeetingNumber field to given value.

### HasPhoneMeetingNumber

`func (o *MeetingBase) HasPhoneMeetingNumber() bool`

HasPhoneMeetingNumber returns a boolean if a field has been set.

### GetVirtualMeetingLink

`func (o *MeetingBase) GetVirtualMeetingLink() string`

GetVirtualMeetingLink returns the VirtualMeetingLink field if non-nil, zero value otherwise.

### GetVirtualMeetingLinkOk

`func (o *MeetingBase) GetVirtualMeetingLinkOk() (*string, bool)`

GetVirtualMeetingLinkOk returns a tuple with the VirtualMeetingLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMeetingLink

`func (o *MeetingBase) SetVirtualMeetingLink(v string)`

SetVirtualMeetingLink sets VirtualMeetingLink field to given value.

### HasVirtualMeetingLink

`func (o *MeetingBase) HasVirtualMeetingLink() bool`

HasVirtualMeetingLink returns a boolean if a field has been set.

### GetVirtualMeetingAdditionalInfo

`func (o *MeetingBase) GetVirtualMeetingAdditionalInfo() string`

GetVirtualMeetingAdditionalInfo returns the VirtualMeetingAdditionalInfo field if non-nil, zero value otherwise.

### GetVirtualMeetingAdditionalInfoOk

`func (o *MeetingBase) GetVirtualMeetingAdditionalInfoOk() (*string, bool)`

GetVirtualMeetingAdditionalInfoOk returns a tuple with the VirtualMeetingAdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMeetingAdditionalInfo

`func (o *MeetingBase) SetVirtualMeetingAdditionalInfo(v string)`

SetVirtualMeetingAdditionalInfo sets VirtualMeetingAdditionalInfo field to given value.

### HasVirtualMeetingAdditionalInfo

`func (o *MeetingBase) HasVirtualMeetingAdditionalInfo() bool`

HasVirtualMeetingAdditionalInfo returns a boolean if a field has been set.

### GetContactName1

`func (o *MeetingBase) GetContactName1() string`

GetContactName1 returns the ContactName1 field if non-nil, zero value otherwise.

### GetContactName1Ok

`func (o *MeetingBase) GetContactName1Ok() (*string, bool)`

GetContactName1Ok returns a tuple with the ContactName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName1

`func (o *MeetingBase) SetContactName1(v string)`

SetContactName1 sets ContactName1 field to given value.

### HasContactName1

`func (o *MeetingBase) HasContactName1() bool`

HasContactName1 returns a boolean if a field has been set.

### GetContactName2

`func (o *MeetingBase) GetContactName2() string`

GetContactName2 returns the ContactName2 field if non-nil, zero value otherwise.

### GetContactName2Ok

`func (o *MeetingBase) GetContactName2Ok() (*string, bool)`

GetContactName2Ok returns a tuple with the ContactName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName2

`func (o *MeetingBase) SetContactName2(v string)`

SetContactName2 sets ContactName2 field to given value.

### HasContactName2

`func (o *MeetingBase) HasContactName2() bool`

HasContactName2 returns a boolean if a field has been set.

### GetContactPhone1

`func (o *MeetingBase) GetContactPhone1() string`

GetContactPhone1 returns the ContactPhone1 field if non-nil, zero value otherwise.

### GetContactPhone1Ok

`func (o *MeetingBase) GetContactPhone1Ok() (*string, bool)`

GetContactPhone1Ok returns a tuple with the ContactPhone1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone1

`func (o *MeetingBase) SetContactPhone1(v string)`

SetContactPhone1 sets ContactPhone1 field to given value.

### HasContactPhone1

`func (o *MeetingBase) HasContactPhone1() bool`

HasContactPhone1 returns a boolean if a field has been set.

### GetContactPhone2

`func (o *MeetingBase) GetContactPhone2() string`

GetContactPhone2 returns the ContactPhone2 field if non-nil, zero value otherwise.

### GetContactPhone2Ok

`func (o *MeetingBase) GetContactPhone2Ok() (*string, bool)`

GetContactPhone2Ok returns a tuple with the ContactPhone2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone2

`func (o *MeetingBase) SetContactPhone2(v string)`

SetContactPhone2 sets ContactPhone2 field to given value.

### HasContactPhone2

`func (o *MeetingBase) HasContactPhone2() bool`

HasContactPhone2 returns a boolean if a field has been set.

### GetContactEmail1

`func (o *MeetingBase) GetContactEmail1() string`

GetContactEmail1 returns the ContactEmail1 field if non-nil, zero value otherwise.

### GetContactEmail1Ok

`func (o *MeetingBase) GetContactEmail1Ok() (*string, bool)`

GetContactEmail1Ok returns a tuple with the ContactEmail1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail1

`func (o *MeetingBase) SetContactEmail1(v string)`

SetContactEmail1 sets ContactEmail1 field to given value.

### HasContactEmail1

`func (o *MeetingBase) HasContactEmail1() bool`

HasContactEmail1 returns a boolean if a field has been set.

### GetContactEmail2

`func (o *MeetingBase) GetContactEmail2() string`

GetContactEmail2 returns the ContactEmail2 field if non-nil, zero value otherwise.

### GetContactEmail2Ok

`func (o *MeetingBase) GetContactEmail2Ok() (*string, bool)`

GetContactEmail2Ok returns a tuple with the ContactEmail2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail2

`func (o *MeetingBase) SetContactEmail2(v string)`

SetContactEmail2 sets ContactEmail2 field to given value.

### HasContactEmail2

`func (o *MeetingBase) HasContactEmail2() bool`

HasContactEmail2 returns a boolean if a field has been set.

### GetBusLines

`func (o *MeetingBase) GetBusLines() string`

GetBusLines returns the BusLines field if non-nil, zero value otherwise.

### GetBusLinesOk

`func (o *MeetingBase) GetBusLinesOk() (*string, bool)`

GetBusLinesOk returns a tuple with the BusLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusLines

`func (o *MeetingBase) SetBusLines(v string)`

SetBusLines sets BusLines field to given value.

### HasBusLines

`func (o *MeetingBase) HasBusLines() bool`

HasBusLines returns a boolean if a field has been set.

### GetTrainLine

`func (o *MeetingBase) GetTrainLine() string`

GetTrainLine returns the TrainLine field if non-nil, zero value otherwise.

### GetTrainLineOk

`func (o *MeetingBase) GetTrainLineOk() (*string, bool)`

GetTrainLineOk returns a tuple with the TrainLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainLine

`func (o *MeetingBase) SetTrainLine(v string)`

SetTrainLine sets TrainLine field to given value.

### HasTrainLine

`func (o *MeetingBase) HasTrainLine() bool`

HasTrainLine returns a boolean if a field has been set.

### GetComments

`func (o *MeetingBase) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *MeetingBase) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *MeetingBase) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *MeetingBase) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


