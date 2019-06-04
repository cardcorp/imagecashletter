/*
 * ImageCashLetter API
 *
 * Moov ImageCashLetter () implements an HTTP API for creating, parsing and validating ImageCashLetter files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

type ReturnDetail struct {
	// CheckDetail ID
	Id string `json:"id,omitempty"`
	// PayorBankRoutingNumber identifies a number that identifies the institution by or through which the item is payable. Must be a valid routing and transit number issued by the ABA’s Routing Number Registrar. Shall represent the first 8 digits of a 9-digit routing number or 8 numeric digits of a 4 dash 4 routing number. A valid routing number consists of 2 fields: the eight- digit Payor Bank Routing Number  and the one-digit Payor Bank Routing Number Check Digit. 
	PayorBankRoutingNumber string `json:"payorBankRoutingNumber,omitempty"`
	// PayorBankCheckDigit identifies a digit representing the routing number check digit.  The combination of Payor Bank Routing Number and payor Bank Routing Number Check Digit must be a mod-checked routing number with a valid check digit. 
	PayorBankCheckDigit string `json:"payorBankCheckDigit,omitempty"`
	// OnUs identifies data specified by the payor bank. On-Us data usually consists of the payor’s account number, a serial number or transaction code, or both.
	OnUs string `json:"onUs,omitempty"`
	// Amount identifies the amount of the check.  All amounts fields have two implied decimal points. e.g., 100000 is $1,000.00
	ItemAmount int32 `json:"itemAmount,omitempty"`
	// ReturnReason is a code that indicates the reason for non-payment.
	ReturnReason string `json:"returnReason,omitempty"`
	// AddendumCount is a number of Check Detail Record Addenda to follow. This represents the number of CheckDetailAddendumA, CheckDetailAddendumB and CheckDetailAddendumC types.  It matches the total number of addendum records associated with this item. The standard supports up to 99 addendum records.
	AddendumCount int32 `json:"addendumCount,omitempty"`
	// DocumentationTypeIndicator identifies a code that indicates the type of documentation that supports the check record.  This field is superseded by the Cash Letter Documentation Type Indicator in the Cash Letter Header Record for all Defined Values except ‘Z’ Not Same Type. In the case of Defined Value of ‘Z’, the Documentation Type Indicator in this record takes precedent.  Shall be present when Cash Letter Documentation Type Indicator in the Cash Letter Header Record is Defined Value of ‘Z’.  * `A` - No image provided, paper provided separately * `B` - No image provided, paper provided separately, image upon request * `C` - Image provided separately, no paper provided * `D` - Image provided separately, no paper provided, image upon request * `E` - Image and paper provided separately * `F` - Image and paper provided separately, image upon request * `G` - Image included, no paper provided * `H` - Image included, no paper provided, image upon request * `I` - Image included, paper provided separately * `J` - Image included, paper provided separately, image upon request * `K` - No image provided, no paper provided * `L` - No image provided, no paper provided, image upon request * `M` - No image provided, Electronic Check provided separately 
	DocumentationTypeIndicator string `json:"documentationTypeIndicator,omitempty"`
	// ForwardBundleDate represents for electronic check exchange items, the year, month, and day that designates the business date of the original forward bundle. This data is transferred from the BundleHeader BundleBusinessDate.  For items presented in paper cash letters, the year, month, and day that the cash letter was created.
	ForwardBundleDate time.Time `json:"forwardBundleDate,omitempty"`
	// ECEInstitutionItemSequenceNumber identifies a number assigned by the institution that creates the CheckDetail. Field must contain a numeric value. It cannot be all blanks.
	ECEInstitutionItemSequenceNumber string `json:"eCEInstitutionItemSequenceNumber,omitempty"`
	// ExternalProcessingCode identifies a code used for special purposes as authorized by the Accredited Standards Committee X9. Also known as Position 44.
	ExternalProcessingCode string `json:"externalProcessingCode,omitempty"`
	// ReturnNotificationIndicator is a code that identifies the type of notification. The CashLetterHeader.CollectionTypeIndicator and the BundleHeader.CollectionTypeIndicator when equal 05 or 06 takes precedence over this field.  * `1` - Preliminary notification * `2` - Final notification 
	ReturnNotificationIndicator string `json:"returnNotificationIndicator,omitempty"`
	// ArchiveTypeIndicator is a code that indicates the type of archive that supports this CheckDetail. Access method, availability and time-frames shall be defined by clearing arrangements. * `A` - Microfilm * `B` - Image * `C` - Paper * `D` - Microfilm and image * `E` - Microfilm and paper * `F` - Image and paper * `G` - Microfilm, image and paper * `H` - Electronic Check Instrument * `I` - None 
	ArchiveTypeIndicator string `json:"archiveTypeIndicator,omitempty"`
	// TimesReturned is code used to indicate the number of times the paying bank has returned this item.  * `0` - The item has been returned an unknown number of times * `1` - The item has been returned once * `2` - The item has been returned twice * `3` - The item has been returned three times 
	TimesReturned int32 `json:"timesReturned,omitempty"`
	ReturnDetailAddendumA ReturnDetailAddendumA `json:"returnDetailAddendumA,omitempty"`
	ReturnDetailAddendumB ReturnDetailAddendumB `json:"returnDetailAddendumB,omitempty"`
	ReturnDetailAddendumC ReturnDetailAddendumC `json:"returnDetailAddendumC,omitempty"`
	ReturnDetailAddendumD ReturnDetailAddendumD `json:"returnDetailAddendumD,omitempty"`
	ImageViewDetail ImageViewDetail `json:"imageViewDetail,omitempty"`
	ImageViewData ImageViewData `json:"imageViewData,omitempty"`
	ImageViewAnalysis ImageViewAnalysis `json:"imageViewAnalysis,omitempty"`
}