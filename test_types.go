package pbsql

type Event struct {
	// @inject_tag: db:"id" primary_key:"y"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" primary_key:"y"`
	// @inject_tag: db:"name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" db:"name"`
	// @inject_tag: db:"description" nullable:"y"
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" db:"description" nullable:"y"`
	// @inject_tag: db:"date_started" nullable:"y"
	DateStarted string `protobuf:"bytes,4,opt,name=date_started,json=dateStarted,proto3" json:"date_started,omitempty" db:"date_started" nullable:"y"`
	// @inject_tag: db:"date_ended" nullable:"y"
	DateEnded string `protobuf:"bytes,5,opt,name=date_ended,json=dateEnded,proto3" json:"date_ended,omitempty" db:"date_ended" nullable:"y"`
	// @inject_tag: db:"time_started" nullable:"y"
	TimeStarted string `protobuf:"bytes,6,opt,name=time_started,json=timeStarted,proto3" json:"time_started,omitempty" db:"time_started" nullable:"y"`
	// @inject_tag: db:"time_ended" nullable:"y"
	TimeEnded string `protobuf:"bytes,7,opt,name=time_ended,json=timeEnded,proto3" json:"time_ended,omitempty" db:"time_ended" nullable:"y"`
	// @inject_tag: db:"is_all_day"
	IsAllDay int32 `protobuf:"varint,8,opt,name=is_all_day,json=isAllDay,proto3" json:"is_all_day,omitempty" db:"is_all_day"`
	// @inject_tag: db:"repeat_type"
	RepeatType int32 `protobuf:"varint,9,opt,name=repeat_type,json=repeatType,proto3" json:"repeat_type,omitempty" db:"repeat_type"`
	// @inject_tag: db:"color"
	Color string `protobuf:"bytes,10,opt,name=color,proto3" json:"color,omitempty" db:"color"`
	// @inject_tag: db:"date_updated"
	DateUpdated string `protobuf:"bytes,11,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty" db:"date_updated"`
	// @inject_tag: db:"date_created"
	DateCreated string `protobuf:"bytes,12,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty" db:"date_created"`
	// @inject_tag: db:"property_id"
	PropertyId int32 `protobuf:"varint,13,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty" db:"property_id"`
	// @inject_tag: db:"contract_id" nullable:"y"
	ContractId int32 `protobuf:"varint,14,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty" db:"contract_id" nullable:"y"`
	// @inject_tag: db:"contract_number" nullable:"y"
	ContractNumber string `protobuf:"bytes,15,opt,name=contract_number,json=contractNumber,proto3" json:"contract_number,omitempty" db:"contract_number" nullable:"y"`
	// @inject_tag: db:"log_jobNumber" nullable:"y"
	LogJobNumber string `protobuf:"bytes,16,opt,name=log_job_number,json=logJobNumber,proto3" json:"log_job_number,omitempty" db:"log_jobNumber" nullable:"y"`
	// @inject_tag: db:"log_jobStatus" nullable:"y"
	LogJobStatus string `protobuf:"bytes,17,opt,name=log_job_status,json=logJobStatus,proto3" json:"log_job_status,omitempty" db:"log_jobStatus" nullable:"y"`
	// @inject_tag: db:"log_PO" nullable:"y"
	LogPo string `protobuf:"bytes,18,opt,name=log_po,json=logPo,proto3" json:"log_po,omitempty" db:"log_PO" nullable:"y"`
	// @inject_tag: db:"log_notes" nullable:"y"
	LogNotes string `protobuf:"bytes,19,opt,name=log_notes,json=logNotes,proto3" json:"log_notes,omitempty" db:"log_notes" nullable:"y"`
	// @inject_tag: db:"log_technicianAssigned" nullable:"y"
	LogTechnicianAssigned string `protobuf:"bytes,20,opt,name=log_technician_assigned,json=logTechnicianAssigned,proto3" json:"log_technician_assigned,omitempty" db:"log_technicianAssigned" nullable:"y"`
	// @inject_tag: db:"log_dateCompleted" nullable:"y"
	LogDateCompleted string `protobuf:"bytes,21,opt,name=log_date_completed,json=logDateCompleted,proto3" json:"log_date_completed,omitempty" db:"log_dateCompleted" nullable:"y"`
	// @inject_tag: db:"log_materialsUsed" nullable:"y"
	LogMaterialsUsed string `protobuf:"bytes,22,opt,name=log_materials_used,json=logMaterialsUsed,proto3" json:"log_materials_used,omitempty" db:"log_materialsUsed" nullable:"y"`
	// @inject_tag: db:"log_serviceRendered" nullable:"y"
	LogServiceRendered string `protobuf:"bytes,23,opt,name=log_service_rendered,json=logServiceRendered,proto3" json:"log_service_rendered,omitempty" db:"log_serviceRendered" nullable:"y"`
	// @inject_tag: db:"log_techNotes" nullable:"y"
	LogTechNotes string `protobuf:"bytes,24,opt,name=log_tech_notes,json=logTechNotes,proto3" json:"log_tech_notes,omitempty" db:"log_techNotes" nullable:"y"`
	// @inject_tag: db:"log_billingDate" nullable:"y"
	LogBillingDate string `protobuf:"bytes,25,opt,name=log_billing_date,json=logBillingDate,proto3" json:"log_billing_date,omitempty" db:"log_billingDate" nullable:"y"`
	// @inject_tag: db:"log_amountCharged" nullable:"y"
	LogAmountCharged string `protobuf:"bytes,26,opt,name=log_amount_charged,json=logAmountCharged,proto3" json:"log_amount_charged,omitempty" db:"log_amountCharged" nullable:"y"`
	// @inject_tag: db:"log_payment_type" nullable:"y"
	LogPaymentType string `protobuf:"bytes,27,opt,name=log_payment_type,json=logPaymentType,proto3" json:"log_payment_type,omitempty" db:"log_payment_type" nullable:"y"`
	// @inject_tag: db:"log_paymentStatus" nullable:"y"
	LogPaymentStatus string `protobuf:"bytes,28,opt,name=log_payment_status,json=logPaymentStatus,proto3" json:"log_payment_status,omitempty" db:"log_paymentStatus" nullable:"y"`
	// @inject_tag: db:"log_timeIn" nullable:"y"
	LogTimeIn string `protobuf:"bytes,29,opt,name=log_time_in,json=logTimeIn,proto3" json:"log_time_in,omitempty" db:"log_timeIn" nullable:"y"`
	// @inject_tag: db:"log_timeOut" nullable:"y"
	LogTimeOut string `protobuf:"bytes,30,opt,name=log_time_out,json=logTimeOut,proto3" json:"log_time_out,omitempty" db:"log_timeOut" nullable:"y"`
	// @inject_tag: db:"log_type" nullable:"y"
	LogType string `protobuf:"bytes,31,opt,name=log_type,json=logType,proto3" json:"log_type,omitempty" db:"log_type" nullable:"y"`
	// @inject_tag: db:"log_contractNotes" nullable:"y"
	LogContractNotes string `protobuf:"bytes,32,opt,name=log_contract_notes,json=logContractNotes,proto3" json:"log_contract_notes,omitempty" db:"log_contractNotes" nullable:"y"`
	// @inject_tag: db:"invoice_serviceItem" nullable:"y"
	InvoiceServiceItem string `protobuf:"bytes,33,opt,name=invoice_service_item,json=invoiceServiceItem,proto3" json:"invoice_service_item,omitempty" db:"invoice_serviceItem" nullable:"y"`
	// @inject_tag: db:"tstat_type" nullable:"y"
	TstatType string `protobuf:"bytes,34,opt,name=tstat_type,json=tstatType,proto3" json:"tstat_type,omitempty" db:"tstat_type" nullable:"y"`
	// @inject_tag: db:"tstat_brand" nullable:"y"
	TstatBrand string `protobuf:"bytes,35,opt,name=tstat_brand,json=tstatBrand,proto3" json:"tstat_brand,omitempty" db:"tstat_brand" nullable:"y"`
	// @inject_tag: db:"compressor_amps" nullable:"y"
	CompressorAmps string `protobuf:"bytes,36,opt,name=compressor_amps,json=compressorAmps,proto3" json:"compressor_amps,omitempty" db:"compressor_amps" nullable:"y"`
	// @inject_tag: db:"condensing_fan_amps" nullable:"y"
	CondensingFanAmps string `protobuf:"bytes,37,opt,name=condensing_fan_amps,json=condensingFanAmps,proto3" json:"condensing_fan_amps,omitempty" db:"condensing_fan_amps" nullable:"y"`
	// @inject_tag: db:"suction_pressure" nullable:"y"
	SuctionPressure string `protobuf:"bytes,38,opt,name=suction_pressure,json=suctionPressure,proto3" json:"suction_pressure,omitempty" db:"suction_pressure" nullable:"y"`
	// @inject_tag: db:"head_pressure" nullable:"y"
	HeadPressure string `protobuf:"bytes,39,opt,name=head_pressure,json=headPressure,proto3" json:"head_pressure,omitempty" db:"head_pressure" nullable:"y"`
	// @inject_tag: db:"return_temperature" nullable:"y"
	ReturnTemperature string `protobuf:"bytes,40,opt,name=return_temperature,json=returnTemperature,proto3" json:"return_temperature,omitempty" db:"return_temperature" nullable:"y"`
	// @inject_tag: db:"supply_temperature" nullable:"y"
	SupplyTemperature string `protobuf:"bytes,41,opt,name=supply_temperature,json=supplyTemperature,proto3" json:"supply_temperature,omitempty" db:"supply_temperature" nullable:"y"`
	// @inject_tag: db:"subcool" nullable:"y"
	Subcool string `protobuf:"bytes,42,opt,name=subcool,proto3" json:"subcool,omitempty" db:"subcool" nullable:"y"`
	// @inject_tag: db:"superheat" nullable:"y"
	Superheat string `protobuf:"bytes,43,opt,name=superheat,proto3" json:"superheat,omitempty" db:"superheat" nullable:"y"`
	// @inject_tag: db:"notes" nullable:"y"
	Notes string `protobuf:"bytes,44,opt,name=notes,proto3" json:"notes,omitempty" db:"notes" nullable:"y"`
	// @inject_tag: db:"services" nullable:"y"
	Services string `protobuf:"bytes,45,opt,name=services,proto3" json:"services,omitempty" db:"services" nullable:"y"`
	// @inject_tag: db:"servicesperformedrow1" nullable:"y"
	Servicesperformedrow1 string `protobuf:"bytes,46,opt,name=servicesperformedrow1,proto3" json:"servicesperformedrow1,omitempty" db:"servicesperformedrow1" nullable:"y"`
	// @inject_tag: db:"totalamountrow1" nullable:"y"
	Totalamountrow1 string `protobuf:"bytes,47,opt,name=totalamountrow1,proto3" json:"totalamountrow1,omitempty" db:"totalamountrow1" nullable:"y"`
	// @inject_tag: db:"servicesperformedrow2" nullable:"y"
	Servicesperformedrow2 string `protobuf:"bytes,48,opt,name=servicesperformedrow2,proto3" json:"servicesperformedrow2,omitempty" db:"servicesperformedrow2" nullable:"y"`
	// @inject_tag: db:"totalamountrow2" nullable:"y"
	Totalamountrow2 string `protobuf:"bytes,49,opt,name=totalamountrow2,proto3" json:"totalamountrow2,omitempty" db:"totalamountrow2" nullable:"y"`
	// @inject_tag: db:"servicesperformedrow3" nullable:"y"
	Servicesperformedrow3 string `protobuf:"bytes,50,opt,name=servicesperformedrow3,proto3" json:"servicesperformedrow3,omitempty" db:"servicesperformedrow3" nullable:"y"`
	// @inject_tag: db:"totalamountrow3" nullable:"y"
	Totalamountrow3 string `protobuf:"bytes,51,opt,name=totalamountrow3,proto3" json:"totalamountrow3,omitempty" db:"totalamountrow3" nullable:"y"`
	// @inject_tag: db:"servicesperformedrow4" nullable:"y"
	Servicesperformedrow4 string `protobuf:"bytes,52,opt,name=servicesperformedrow4,proto3" json:"servicesperformedrow4,omitempty" db:"servicesperformedrow4" nullable:"y"`
	// @inject_tag: db:"totalamountrow4" nullable:"y"
	Totalamountrow4 string `protobuf:"bytes,53,opt,name=totalamountrow4,proto3" json:"totalamountrow4,omitempty" db:"totalamountrow4" nullable:"y"`
	// @inject_tag: db:"discount" nullable:"y"
	Discount string `protobuf:"bytes,54,opt,name=discount,proto3" json:"discount,omitempty" db:"discount" nullable:"y"`
	// @inject_tag: db:"discountcost" nullable:"y"
	Discountcost string `protobuf:"bytes,55,opt,name=discountcost,proto3" json:"discountcost,omitempty" db:"discountcost" nullable:"y"`
	// @inject_tag: db:"log_notification" nullable:"y"
	LogNotification string `protobuf:"bytes,56,opt,name=log_notification,json=logNotification,proto3" json:"log_notification,omitempty" db:"log_notification" nullable:"y"`
	// @inject_tag: db:"diagnosticQuoted"
	DiagnosticQuoted int32 `protobuf:"varint,57,opt,name=diagnostic_quoted,json=diagnosticQuoted,proto3" json:"diagnostic_quoted,omitempty" db:"diagnosticQuoted"`
	// @inject_tag: db:"amountQuoted" nullable:"y"
	AmountQuoted string `protobuf:"bytes,58,opt,name=amount_quoted,json=amountQuoted,proto3" json:"amount_quoted,omitempty" db:"amountQuoted" nullable:"y"`
	// @inject_tag: db:"propertyBilling"
	PropertyBilling int32 `protobuf:"varint,59,opt,name=property_billing,json=propertyBilling,proto3" json:"property_billing,omitempty" db:"propertyBilling"`
	// @inject_tag: db:"isCallback" nullable:"y"
	IsCallback int32 `protobuf:"varint,60,opt,name=is_callback,json=isCallback,proto3" json:"is_callback,omitempty" db:"isCallback" nullable:"y"`
	// @inject_tag: db:"log_version"
	LogVersion int32 `protobuf:"varint,61,opt,name=log_version,json=logVersion,proto3" json:"log_version,omitempty" db:"log_version"`
	// @inject_tag: db:"job_type_id" nullable:"y"
	JobTypeId int32 `protobuf:"varint,62,opt,name=job_type_id,json=jobTypeId,proto3" json:"job_type_id,omitempty" db:"job_type_id" nullable:"y"`
	// @inject_tag: db:"job_subtype_id" nullable:"y"
	JobSubtypeId int32 `protobuf:"varint,63,opt,name=job_subtype_id,json=jobSubtypeId,proto3" json:"job_subtype_id,omitempty" db:"job_subtype_id" nullable:"y"`
	// @inject_tag: db:"callback_original_id" nullable:"y"
	CallbackOriginalId int32 `protobuf:"varint,64,opt,name=callback_original_id,json=callbackOriginalId,proto3" json:"callback_original_id,omitempty" db:"callback_original_id" nullable:"y"`
	// @inject_tag: db:"callback_disposition" nullable:"y"
	CallbackDisposition int32 `protobuf:"varint,65,opt,name=callback_disposition,json=callbackDisposition,proto3" json:"callback_disposition,omitempty" db:"callback_disposition" nullable:"y"`
	// @inject_tag: db:"callback_comments" nullable:"y"
	CallbackComments string `protobuf:"bytes,66,opt,name=callback_comments,json=callbackComments,proto3" json:"callback_comments,omitempty" db:"callback_comments" nullable:"y"`
	// @inject_tag: db:"callback_technician" nullable:"y"
	CallbackTechnician int32 `protobuf:"varint,67,opt,name=callback_technician,json=callbackTechnician,proto3" json:"callback_technician,omitempty" db:"callback_technician" nullable:"y"`
	// @inject_tag: db:"callback_approval_timestamp" nullable:"y"
	CallbackApprovalTimestamp string `protobuf:"bytes,68,opt,name=callback_approval_timestamp,json=callbackApprovalTimestamp,proto3" json:"callback_approval_timestamp,omitempty" db:"callback_approval_timestamp" nullable:"y"`
	// @inject_tag: db:"callback_comment_by" nullable:"y"
	CallbackCommentBy int32 `protobuf:"varint,69,opt,name=callback_comment_by,json=callbackCommentBy,proto3" json:"callback_comment_by,omitempty" db:"callback_comment_by" nullable:"y"`
	// @inject_tag: db:"document_id" nullable:"y"
	DocumentId int32 `protobuf:"varint,70,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty" db:"document_id" nullable:"y"`
	// @inject_tag: db:"material_used" nullable:"y"
	MaterialUsed string `protobuf:"bytes,71,opt,name=material_used,json=materialUsed,proto3" json:"material_used,omitempty" db:"material_used" nullable:"y"`
	// @inject_tag: db:"material_total" nullable:"y"
	MaterialTotal float64 `protobuf:"fixed64,72,opt,name=material_total,json=materialTotal,proto3" json:"material_total,omitempty" db:"material_total" nullable:"y"`
	// @inject_tag: db:"isActive" nullable:"y"
	IsActive int32 `protobuf:"varint,73,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty" db:"isActive" nullable:"y"`
	// @inject_tag: db:"parent_id" nullable:"y"
	ParentId int32 `protobuf:"varint,74,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty" db:"parent_id" nullable:"y"`
	// @inject_tag: db:"isLmpc"
	IsLmpc int32 `protobuf:"varint,75,opt,name=is_lmpc,json=isLmpc,proto3" json:"is_lmpc,omitempty" db:"isLmpc"`
	// @inject_tag: db:"highPriority"
	HighPriority int32 `protobuf:"varint,76,opt,name=high_priority,json=highPriority,proto3" json:"high_priority,omitempty" db:"highPriority"`
	// @inject_tag: db:"isResidential"
	IsResidential int32      `protobuf:"varint,77,opt,name=is_residential,json=isResidential,proto3" json:"is_residential,omitempty" db:"isResidential"`
	JobType       string     `protobuf:"bytes,78,opt,name=job_type,json=jobType,proto3" json:"job_type,omitempty"`
	JobSubtype    string     `protobuf:"bytes,79,opt,name=job_subtype,json=jobSubtype,proto3" json:"job_subtype,omitempty"`
	Customer      *User `protobuf:"bytes,80,opt,name=customer,proto3" json:"customer,omitempty" local_name:"id" foreign_key:"id" foreign_table:"servicable2"`
	// @inject_tag: foreign_key:"property_id" foreign_table:"properties"
	Property             *Property `protobuf:"bytes,81,opt,name=property,proto3" json:"property,omitempty" local_name:"property_id" foreign_key:"property_id" foreign_table:"properties"`
	FieldMask            []string           `protobuf:"bytes,82,rep,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	PageNumber           int32              `protobuf:"varint,83,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	OrderBy              string             `protobuf:"bytes,84,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	OrderDir             string             `protobuf:"bytes,85,opt,name=order_dir,json=orderDir,proto3" json:"order_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
	DateRange 					 []string           `date_target:"date_started"`
}

type Property struct {
	// @inject_tag: db:"property_id" primary_key:"y"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"property_id" primary_key:"y"`
	// @inject_tag: db:"user_id" nullable:"y"
	UserId int32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" db:"user_id" nullable:"y"`
	// @inject_tag: db:"contract_id" nullable:"y"
	ContractId int32 `protobuf:"varint,3,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty" db:"contract_id" nullable:"y"`
	// @inject_tag: db:"property_address" nullable:"y"
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty" db:"property_address" nullable:"y"`
	// @inject_tag: db:"property_city" nullable:"y"
	City string `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty" db:"property_city" nullable:"y"`
	// @inject_tag: db:"property_state" nullable:"y"
	State string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty" db:"property_state" nullable:"y"`
	// @inject_tag: db:"property_zip" nullable:"y"
	Zip string `protobuf:"bytes,7,opt,name=zip,proto3" json:"zip,omitempty" db:"property_zip" nullable:"y"`
	// @inject_tag: db:"property_subdivision" nullable:"y"
	Subdivision string `protobuf:"bytes,8,opt,name=subdivision,proto3" json:"subdivision,omitempty" db:"property_subdivision" nullable:"y"`
	// @inject_tag: db:"property_directions" nullable:"y"
	Directions string `protobuf:"bytes,9,opt,name=directions,proto3" json:"directions,omitempty" db:"property_directions" nullable:"y"`
	// @inject_tag: db:"property_notes" nullable:"y"
	Notes string `protobuf:"bytes,10,opt,name=notes,proto3" json:"notes,omitempty" db:"property_notes" nullable:"y"`
	// @inject_tag: db:"property_date_created" nullable:"y"
	DateCreated string `protobuf:"bytes,11,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty" db:"property_date_created" nullable:"y"`
	// @inject_tag: db:"property_isActive"
	IsActive int32 `protobuf:"varint,12,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty" db:"property_isActive"`
	// @inject_tag: db:"property_isResidential" nullable:"y"
	IsResidential int32 `protobuf:"varint,13,opt,name=is_residential,json=isResidential,proto3" json:"is_residential,omitempty" db:"property_isResidential" nullable:"y"`
	// @inject_tag: db:"property_notification" nullable:"y"
	Notification string `protobuf:"bytes,14,opt,name=notification,proto3" json:"notification,omitempty" db:"property_notification" nullable:"y"`
	// @inject_tag: db:"property_firstname" nullable:"y"
	Firstname string `protobuf:"bytes,15,opt,name=firstname,proto3" json:"firstname,omitempty" db:"property_firstname" nullable:"y"`
	// @inject_tag: db:"property_lastname" nullable:"y"
	Lastname string `protobuf:"bytes,16,opt,name=lastname,proto3" json:"lastname,omitempty" db:"property_lastname" nullable:"y"`
	// @inject_tag: db:"property_businessname" nullable:"y"
	Businessname string `protobuf:"bytes,17,opt,name=businessname,proto3" json:"businessname,omitempty" db:"property_businessname" nullable:"y"`
	// @inject_tag: db:"property_phone" nullable:"y"
	Phone string `protobuf:"bytes,18,opt,name=phone,proto3" json:"phone,omitempty" db:"property_phone" nullable:"y"`
	// @inject_tag: db:"property_altphone" nullable:"y"
	Altphone string `protobuf:"bytes,19,opt,name=altphone,proto3" json:"altphone,omitempty" db:"property_altphone" nullable:"y"`
	// @inject_tag: db:"property_email" nullable:"y"
	Email string `protobuf:"bytes,20,opt,name=email,proto3" json:"email,omitempty" db:"property_email" nullable:"y"`
	// @inject_tag: db:"geolocation_lat" nullable:"y"
	GeolocationLat float64 `protobuf:"fixed64,21,opt,name=geolocation_lat,json=geolocationLat,proto3" json:"geolocation_lat,omitempty" db:"geolocation_lat" nullable:"y"`
	// @inject_tag: db:"geolocation_lng" nullable:"y"
	GeolocationLng       float64  `protobuf:"fixed64,22,opt,name=geolocation_lng,json=geolocationLng,proto3" json:"geolocation_lng,omitempty" db:"geolocation_lng" nullable:"y"`
	FieldMask            []string `protobuf:"bytes,23,rep,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	PageNumber           int32    `protobuf:"varint,24,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type User struct {
	// @inject_tag: db:"user_id" primary_key:"y"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"user_id" primary_key:"y"`
	// @inject_tag: db:"user_firstname" nullable:"y"
	Firstname string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty" db:"user_firstname" nullable:"y"`
	// @inject_tag: db:"user_lastname" nullable:"y"
	Lastname string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty" db:"user_lastname" nullable:"y"`
	// @inject_tag: db:"user_businessname" nullable:"y"
	Businessname string `protobuf:"bytes,4,opt,name=businessname,proto3" json:"businessname,omitempty" db:"user_businessname" nullable:"y"`
	// @inject_tag: db:"user_city" nullable:"y"
	City string `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty" db:"user_city" nullable:"y"`
	// @inject_tag: db:"user_state" nullable:"y"
	State string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty" db:"user_state" nullable:"y"`
	// @inject_tag: db:"user_zip" nullable:"y"
	Zip string `protobuf:"bytes,7,opt,name=zip,proto3" json:"zip,omitempty" db:"user_zip" nullable:"y"`
	// @inject_tag: db:"user_address" nullable:"y"
	Address string `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty" db:"user_address" nullable:"y"`
	// @inject_tag: db:"user_phone" nullable:"y"
	Phone string `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty" db:"user_phone" nullable:"y"`
	// @inject_tag: db:"user_altphone" nullable:"y"
	Altphone string `protobuf:"bytes,10,opt,name=altphone,proto3" json:"altphone,omitempty" db:"user_altphone" nullable:"y"`
	// @inject_tag: db:"user_cellphone" nullable:"y"
	Cellphone string `protobuf:"bytes,11,opt,name=cellphone,proto3" json:"cellphone,omitempty" db:"user_cellphone" nullable:"y"`
	// @inject_tag: db:"user_fax" nullable:"y"
	Fax string `protobuf:"bytes,12,opt,name=fax,proto3" json:"fax,omitempty" db:"user_fax" nullable:"y"`
	// @inject_tag: db:"user_email" nullable:"y"
	Email string `protobuf:"bytes,13,opt,name=email,proto3" json:"email,omitempty" db:"user_email" nullable:"y"`
	// @inject_tag: db:"user_alt_email" nullable:"y"
	AltEmail string `protobuf:"bytes,14,opt,name=alt_email,json=altEmail,proto3" json:"alt_email,omitempty" db:"user_alt_email" nullable:"y"`
	// @inject_tag: db:"user_phone_email" nullable:"y"
	PhoneEmail string `protobuf:"bytes,15,opt,name=phone_email,json=phoneEmail,proto3" json:"phone_email,omitempty" db:"user_phone_email" nullable:"y"`
	// @inject_tag: db:"user_preferredContact" nullable:"y"
	PreferredContact string `protobuf:"bytes,16,opt,name=preferred_contact,json=preferredContact,proto3" json:"preferred_contact,omitempty" db:"user_preferredContact" nullable:"y"`
	// @inject_tag: db:"user_receiveemail" nullable:"y"
	Receiveemail int32 `protobuf:"varint,17,opt,name=receiveemail,proto3" json:"receiveemail,omitempty" db:"user_receiveemail" nullable:"y"`
	// @inject_tag: db:"user_date_created" nullable:"y"
	DateCreated string `protobuf:"bytes,18,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty" db:"user_date_created" nullable:"y"`
	// @inject_tag: db:"user_last_login" nullable:"y"
	LastLogin string `protobuf:"bytes,19,opt,name=last_login,json=lastLogin,proto3" json:"last_login,omitempty" db:"user_last_login" nullable:"y"`
	// @inject_tag: db:"annual_hours_pto" nullable:"y"
	AnnualHoursPto float64 `protobuf:"fixed64,21,opt,name=annual_hours_pto,json=annualHoursPto,proto3" json:"annual_hours_pto,omitempty" db:"annual_hours_pto" nullable:"y"`
	// @inject_tag: db:"bonus_hours_pto" nullable:"y"
	BonusHoursPto float64 `protobuf:"fixed64,22,opt,name=bonus_hours_pto,json=bonusHoursPto,proto3" json:"bonus_hours_pto,omitempty" db:"bonus_hours_pto" nullable:"y"`
	// @inject_tag: db:"user_isActive" nullable:"y"
	IsActive int32 `protobuf:"varint,23,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty" db:"user_isActive" nullable:"y"`
	// @inject_tag: db:"user_isSU" nullable:"y"
	Is_SU int32 `protobuf:"varint,24,opt,name=is_SU,json=isSU,proto3" json:"is_SU,omitempty" db:"user_isSU" nullable:"y"`
	// @inject_tag: db:"user_isAdmin" nullable:"y"
	IsAdmin int32 `protobuf:"varint,25,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty" db:"user_isAdmin" nullable:"y"`
	// @inject_tag: db:"is_office_staff" nullable:"y"
	IsOfficeStaff int32 `protobuf:"varint,26,opt,name=is_office_staff,json=isOfficeStaff,proto3" json:"is_office_staff,omitempty" db:"is_office_staff" nullable:"y"`
	// @inject_tag: db:"user_office_group" nullable:"y"
	OfficeGroup int32 `protobuf:"varint,27,opt,name=office_group,json=officeGroup,proto3" json:"office_group,omitempty" db:"user_office_group" nullable:"y"`
	// @inject_tag: db:"is_hvac_tech" nullable:"y"
	IsHvacTech int32 `protobuf:"varint,28,opt,name=is_hvac_tech,json=isHvacTech,proto3" json:"is_hvac_tech,omitempty" db:"is_hvac_tech" nullable:"y"`
	// @inject_tag: db:"tech_assist" nullable:"y"
	TechAssist int32 `protobuf:"varint,29,opt,name=tech_assist,json=techAssist,proto3" json:"tech_assist,omitempty" db:"tech_assist" nullable:"y"`
	// @inject_tag: db:"user_calendarPref" nullable:"y"
	CalendarPref string `protobuf:"bytes,30,opt,name=calendar_pref,json=calendarPref,proto3" json:"calendar_pref,omitempty" db:"user_calendarPref" nullable:"y"`
	// @inject_tag: db:"user_multiProperty" nullable:"y"
	MultiProperty int32 `protobuf:"varint,31,opt,name=multi_property,json=multiProperty,proto3" json:"multi_property,omitempty" db:"user_multiProperty" nullable:"y"`
	// @inject_tag: db:"user_isEmployee" nullable:"y"
	IsEmployee int32 `protobuf:"varint,32,opt,name=is_employee,json=isEmployee,proto3" json:"is_employee,omitempty" db:"user_isEmployee" nullable:"y"`
	// @inject_tag: db:"employee_function_id" nullable:"y"
	EmployeeFunctionId int32 `protobuf:"varint,33,opt,name=employee_function_id,json=employeeFunctionId,proto3" json:"employee_function_id,omitempty" db:"employee_function_id" nullable:"y"`
	// @inject_tag: db:"employee_department_id" nullable:"y"
	EmployeeDepartmentId int32 `protobuf:"varint,34,opt,name=employee_department_id,json=employeeDepartmentId,proto3" json:"employee_department_id,omitempty" db:"employee_department_id" nullable:"y"`
	// @inject_tag: db:"user_login" nullable:"y"
	Login string `protobuf:"bytes,35,opt,name=login,proto3" json:"login,omitempty" db:"user_login" nullable:"y"`
	// @inject_tag: db:"user_pwd" nullable:"y"
	Pwd string `protobuf:"bytes,36,opt,name=pwd,proto3" json:"pwd,omitempty" db:"user_pwd" nullable:"y"`
	// @inject_tag: db:"user_notes" nullable:"y"
	Notes string `protobuf:"bytes,37,opt,name=notes,proto3" json:"notes,omitempty" db:"user_notes" nullable:"y"`
	// @inject_tag: db:"user_intNotes" nullable:"y"
	IntNotes string `protobuf:"bytes,38,opt,name=int_notes,json=intNotes,proto3" json:"int_notes,omitempty" db:"user_intNotes" nullable:"y"`
	// @inject_tag: db:"user_notification" nullable:"y"
	Notification string `protobuf:"bytes,39,opt,name=notification,proto3" json:"notification,omitempty" db:"user_notification" nullable:"y"`
	// @inject_tag: db:"user_billingTerms" nullable:"y"
	BillingTerms string `protobuf:"bytes,40,opt,name=billing_terms,json=billingTerms,proto3" json:"billing_terms,omitempty" db:"user_billingTerms" nullable:"y"`
	// @inject_tag: db:"user_rebate" nullable:"y"
	Rebate int32 `protobuf:"varint,41,opt,name=rebate,proto3" json:"rebate,omitempty" db:"user_rebate" nullable:"y"`
	// @inject_tag: db:"user_discount" nullable:"y"
	Discount int32 `protobuf:"varint,42,opt,name=discount,proto3" json:"discount,omitempty" db:"user_discount" nullable:"y"`
	// @inject_tag: db:"user_managed_by" nullable:"y"
	ManagedBy int32 `protobuf:"varint,43,opt,name=managed_by,json=managedBy,proto3" json:"managed_by,omitempty" db:"user_managed_by" nullable:"y"`
	// @inject_tag: db:"current_status" nullable:"y"
	CurrentStatus string `protobuf:"bytes,44,opt,name=current_status,json=currentStatus,proto3" json:"current_status,omitempty" db:"current_status" nullable:"y"`
	// @inject_tag: db:"current_status_jobNumber" nullable:"y"
	CurrentStatusJobNumber string `protobuf:"bytes,45,opt,name=current_status_job_number,json=currentStatusJobNumber,proto3" json:"current_status_job_number,omitempty" db:"current_status_jobNumber" nullable:"y"`
	// @inject_tag: db:"current_status_timestamp" nullable:"y"
	CurrentStatusTimestamp string `protobuf:"bytes,46,opt,name=current_status_timestamp,json=currentStatusTimestamp,proto3" json:"current_status_timestamp,omitempty" db:"current_status_timestamp" nullable:"y"`
	// @inject_tag: db:"emp_title" nullable:"y"
	EmpTitle string `protobuf:"bytes,47,opt,name=emp_title,json=empTitle,proto3" json:"emp_title,omitempty" db:"emp_title" nullable:"y"`
	// @inject_tag: db:"ext" nullable:"y"
	Ext string `protobuf:"bytes,48,opt,name=ext,proto3" json:"ext,omitempty" db:"ext" nullable:"y"`
	// @inject_tag: db:"image" nullable:"y"
	Image string `protobuf:"bytes,49,opt,name=image,proto3" json:"image,omitempty" db:"image" nullable:"y"`
	// @inject_tag: db:"user_serviceCalls" nullable:"y"
	ServiceCalls int32 `protobuf:"varint,50,opt,name=service_calls,json=serviceCalls,proto3" json:"service_calls,omitempty" db:"user_serviceCalls" nullable:"y"`
	// @inject_tag: db:"user_show_billing" nullable:"y"
	ShowBilling int32 `protobuf:"varint,51,opt,name=show_billing,json=showBilling,proto3" json:"show_billing,omitempty" db:"user_show_billing" nullable:"y"`
	// @inject_tag: db:"paid_service_call_status" nullable:"y"
	PaidServiceCallStatus int32 `protobuf:"varint,52,opt,name=paid_service_call_status,json=paidServiceCallStatus,proto3" json:"paid_service_call_status,omitempty" db:"paid_service_call_status" nullable:"y"`
	// @inject_tag: db:"is_color_mute" nullable:"y"
	IsColorMute int32 `protobuf:"varint,53,opt,name=is_color_mute,json=isColorMute,proto3" json:"is_color_mute,omitempty" db:"is_color_mute" nullable:"y"`
	// @inject_tag: db:"service_call_refresh" nullable:"y"
	ServiceCallRefresh int32 `protobuf:"varint,54,opt,name=service_call_refresh,json=serviceCallRefresh,proto3" json:"service_call_refresh,omitempty" db:"service_call_refresh" nullable:"y"`
	// @inject_tag: db:"tool_fund" nullable:"y"
	ToolFund float64 `protobuf:"fixed64,55,opt,name=tool_fund,json=toolFund,proto3" json:"tool_fund,omitempty" db:"tool_fund" nullable:"y"`
	// @inject_tag: db:"spiff_fund" nullable:"y"
	SpiffFund float64 `protobuf:"fixed64,56,opt,name=spiff_fund,json=spiffFund,proto3" json:"spiff_fund,omitempty" db:"spiff_fund" nullable:"y"`
	// @inject_tag: db:"geolocation_lat" nullable:"y"
	GeolocationLat float64 `protobuf:"fixed64,57,opt,name=geolocation_lat,json=geolocationLat,proto3" json:"geolocation_lat,omitempty" db:"geolocation_lat" nullable:"y"`
	// @inject_tag: db:"geolocation_lng" nullable:"y"
	GeolocationLng float64 `protobuf:"fixed64,58,opt,name=geolocation_lng,json=geolocationLng,proto3" json:"geolocation_lng,omitempty" db:"geolocation_lng" nullable:"y"`
	// @inject_tag: foreign_key:"technician_user_id" foreign_table:"services_rendered" local_name:"user_id"
	ServicesRendered     *ServicesRendered `protobuf:"bytes,63,opt,name=services_rendered,json=servicesRendered,proto3" json:"services_rendered,omitempty" foreign_key:"technician_user_id" foreign_table:"services_rendered" local_name:"user_id"`
	FieldMask            []string                            `protobuf:"bytes,59,rep,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	PageNumber           int32                               `protobuf:"varint,60,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	OrderBy              string                              `protobuf:"bytes,61,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	OrderDir             string                              `protobuf:"bytes,62,opt,name=order_dir,json=orderDir,proto3" json:"order_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

type ServicesRendered struct {
	// @inject_tag: db:"sr_id" primary_key:"y"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"sr_id" primary_key:"y"`
	// @inject_tag: db:"event_id" nullable:"y"
	EventId int32 `protobuf:"varint,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty" db:"event_id" nullable:"y"`
	// @inject_tag: db:"technician_user_id" nullable:"y"
	TechnicianUserId int32 `protobuf:"varint,3,opt,name=technician_user_id,json=technicianUserId,proto3" json:"technician_user_id,omitempty" db:"technician_user_id" nullable:"y"`
	// @inject_tag: db:"sr_name"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" db:"sr_name"`
	// @inject_tag: db:"sr_materialsUsed" nullable:"y"
	MaterialsUsed string `protobuf:"bytes,5,opt,name=materials_used,json=materialsUsed,proto3" json:"materials_used,omitempty" db:"sr_materialsUsed" nullable:"y"`
	// @inject_tag: db:"sr_serviceRendered" nullable:"y"
	ServiceRendered string `protobuf:"bytes,6,opt,name=service_rendered,json=serviceRendered,proto3" json:"service_rendered,omitempty" db:"sr_serviceRendered" nullable:"y"`
	// @inject_tag: db:"sr_techNotes" nullable:"y"
	TechNotes string `protobuf:"bytes,7,opt,name=tech_notes,json=techNotes,proto3" json:"tech_notes,omitempty" db:"sr_techNotes" nullable:"y"`
	// @inject_tag: db:"sr_status"
	Status string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty" db:"sr_status"`
	// @inject_tag: db:"sr_datetime"
	Datetime string `protobuf:"bytes,9,opt,name=datetime,proto3" json:"datetime,omitempty" db:"sr_datetime"`
	// @inject_tag: db:"time_started" nullable:"y"
	TimeStarted string `protobuf:"bytes,10,opt,name=time_started,json=timeStarted,proto3" json:"time_started,omitempty" db:"time_started" nullable:"y"`
	// @inject_tag: db:"time_finished" nullable:"y"
	TimeFinished string `protobuf:"bytes,11,opt,name=time_finished,json=timeFinished,proto3" json:"time_finished,omitempty" db:"time_finished" nullable:"y"`
	// @inject_tag: db:"isactive" nullable:"y"
	IsActive int32 `protobuf:"varint,12,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty" db:"isactive" nullable:"y"`
	// @inject_tag: db:"hide_from_timesheet" nullable:"y"
	HideFromTimesheet int32 `protobuf:"varint,13,opt,name=hide_from_timesheet,json=hideFromTimesheet,proto3" json:"hide_from_timesheet,omitempty" db:"hide_from_timesheet" nullable:"y"`
	// @inject_tag: db:"signature_id" nullable:"y"
	SignatureId int32 `protobuf:"varint,14,opt,name=signature_id,json=signatureId,proto3" json:"signature_id,omitempty" db:"signature_id" nullable:"y"`
	// @inject_tag: db:"signatureData" nullable:"y"
	SignatureData        string   `protobuf:"bytes,15,opt,name=signature_data,json=signatureData,proto3" json:"signature_data,omitempty" db:"signatureData" nullable:"y"`
	FieldMask            []string `protobuf:"bytes,16,rep,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	PageNumber           int32    `protobuf:"varint,17,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type Transaction struct {
	// @inject_tag: db:"id" primary_key:"y"
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" primary_key:"y"`
	// @inject_tag: db:"job_id" nullable:"y"
	JobId int32 `protobuf:"varint,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty" db:"job_id" nullable:"y"`
	// @inject_tag: db:"department_id" nullable:"y"
	DepartmentId int32 `protobuf:"varint,3,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty" db:"department_id" nullable:"y"`
	// @inject_tag: db:"owner_id" nullable:"y"
	OwnerId int32 `protobuf:"varint,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty" db:"owner_id" nullable:"y"`
	// @inject_tag: db:"vendor" nullable:"y"
	Vendor string `protobuf:"bytes,5,opt,name=vendor,proto3" json:"vendor,omitempty" db:"vendor" nullable:"y"`
	// @inject_tag: db:"cost_center_id" nullable:"y"
	CostCenterId int32 `protobuf:"varint,6,opt,name=cost_center_id,json=costCenterId,proto3" json:"cost_center_id,omitempty" db:"cost_center_id" nullable:"y"`
	// @inject_tag: db:"description" nullable:"y"
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty" db:"description" nullable:"y"`
	// @inject_tag: db:"amount" nullable:"y"
	Amount float64 `protobuf:"fixed64,8,opt,name=amount,proto3" json:"amount,omitempty" db:"amount" nullable:"y"`
	// @inject_tag: db:"timestamp" nullable:"y"
	Timestamp string `protobuf:"bytes,9,opt,name=timestamp,proto3" json:"timestamp,omitempty" db:"timestamp" nullable:"y"`
	// @inject_tag: db:"notes" nullable:"y"
	Notes string `protobuf:"bytes,10,opt,name=notes,proto3" json:"notes,omitempty" db:"notes" nullable:"y"`
	// @inject_tag: db:"is_active" nullable:"y"
	IsActive int32 `protobuf:"varint,11,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty" db:"is_active" nullable:"y"`
	// @inject_tag: db:"status_id"
	StatusId int32 `protobuf:"varint,12,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty" db:"status_id"`
	// @inject_tag: db:"is_audited"
	IsAudited bool `protobuf:"varint,24,opt,name=is_audited,json=isAudited,proto3" json:"is_audited,omitempty" db:"is_audited"`
	// @inect_tag: db:"is_recorded"
	IsRecorded           bool                                        `protobuf:"varint,25,opt,name=is_recorded,json=isRecorded,proto3" json:"is_recorded,omitempty" db:"is_recorded"`
	Status               string                                      `protobuf:"bytes,13,opt,name=status,proto3" json:"status,omitempty"`
	OwnerName            string                                      `protobuf:"bytes,14,opt,name=owner_name,json=ownerName,proto3" json:"owner_name,omitempty"`
	CardUsed             string                                      `protobuf:"bytes,15,opt,name=card_used,json=cardUsed,proto3" json:"card_used,omitempty"`
	PageNumber           int32                                       `protobuf:"varint,18,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	FieldMask            []string                                    `protobuf:"bytes,19,rep,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	OrderBy              string                                      `protobuf:"bytes,22,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	OrderDir             string                                      `protobuf:"bytes,23,opt,name=order_dir,json=orderDir,proto3" json:"order_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
	DateRange []string `date_target:"timestamp"`
	DateTarget string
}
