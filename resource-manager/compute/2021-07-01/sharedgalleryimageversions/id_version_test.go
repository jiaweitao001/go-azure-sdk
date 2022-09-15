package sharedgalleryimageversions

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = VersionId{}

func TestNewVersionID(t *testing.T) {
	id := NewVersionID("12345678-1234-9876-4563-123456789012", "locationValue", "galleryUniqueValue", "galleryImageValue", "galleryImageVersionValue")

	if id.SubscriptionId != "12345678-1234-9876-4563-123456789012" {
		t.Fatalf("Expected %q but got %q for Segment 'SubscriptionId'", id.SubscriptionId, "12345678-1234-9876-4563-123456789012")
	}

	if id.Location != "locationValue" {
		t.Fatalf("Expected %q but got %q for Segment 'Location'", id.Location, "locationValue")
	}

	if id.GalleryUniqueName != "galleryUniqueValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GalleryUniqueName'", id.GalleryUniqueName, "galleryUniqueValue")
	}

	if id.GalleryImageName != "galleryImageValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GalleryImageName'", id.GalleryImageName, "galleryImageValue")
	}

	if id.GalleryImageVersionName != "galleryImageVersionValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GalleryImageVersionName'", id.GalleryImageVersionName, "galleryImageVersionValue")
	}
}

func TestFormatVersionID(t *testing.T) {
	actual := NewVersionID("12345678-1234-9876-4563-123456789012", "locationValue", "galleryUniqueValue", "galleryImageValue", "galleryImageVersionValue").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue/images/galleryImageValue/versions/galleryImageVersionValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseVersionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue/images",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue/images/galleryImageValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue/images/galleryImageValue/versions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue/images/galleryImageValue/versions/galleryImageVersionValue",
			Expected: &VersionId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				Location:                "locationValue",
				GalleryUniqueName:       "galleryUniqueValue",
				GalleryImageName:        "galleryImageValue",
				GalleryImageVersionName: "galleryImageVersionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue/images/galleryImageValue/versions/galleryImageVersionValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVersionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.Location != v.Expected.Location {
			t.Fatalf("Expected %q but got %q for Location", v.Expected.Location, actual.Location)
		}

		if actual.GalleryUniqueName != v.Expected.GalleryUniqueName {
			t.Fatalf("Expected %q but got %q for GalleryUniqueName", v.Expected.GalleryUniqueName, actual.GalleryUniqueName)
		}

		if actual.GalleryImageName != v.Expected.GalleryImageName {
			t.Fatalf("Expected %q but got %q for GalleryImageName", v.Expected.GalleryImageName, actual.GalleryImageName)
		}

		if actual.GalleryImageVersionName != v.Expected.GalleryImageVersionName {
			t.Fatalf("Expected %q but got %q for GalleryImageVersionName", v.Expected.GalleryImageVersionName, actual.GalleryImageVersionName)
		}

	}
}

func TestParseVersionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *VersionId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/sHaReDgAlLeRiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/sHaReDgAlLeRiEs/gAlLeRyUnIqUeVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue/images",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/sHaReDgAlLeRiEs/gAlLeRyUnIqUeVaLuE/iMaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue/images/galleryImageValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/sHaReDgAlLeRiEs/gAlLeRyUnIqUeVaLuE/iMaGeS/gAlLeRyImAgEvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue/images/galleryImageValue/versions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/sHaReDgAlLeRiEs/gAlLeRyUnIqUeVaLuE/iMaGeS/gAlLeRyImAgEvAlUe/vErSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue/images/galleryImageValue/versions/galleryImageVersionValue",
			Expected: &VersionId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				Location:                "locationValue",
				GalleryUniqueName:       "galleryUniqueValue",
				GalleryImageName:        "galleryImageValue",
				GalleryImageVersionName: "galleryImageVersionValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/providers/Microsoft.Compute/locations/locationValue/sharedGalleries/galleryUniqueValue/images/galleryImageValue/versions/galleryImageVersionValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/sHaReDgAlLeRiEs/gAlLeRyUnIqUeVaLuE/iMaGeS/gAlLeRyImAgEvAlUe/vErSiOnS/gAlLeRyImAgEvErSiOnVaLuE",
			Expected: &VersionId{
				SubscriptionId:          "12345678-1234-9876-4563-123456789012",
				Location:                "lOcAtIoNvAlUe",
				GalleryUniqueName:       "gAlLeRyUnIqUeVaLuE",
				GalleryImageName:        "gAlLeRyImAgEvAlUe",
				GalleryImageVersionName: "gAlLeRyImAgEvErSiOnVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/sUbScRiPtIoNs/12345678-1234-9876-4563-123456789012/pRoViDeRs/mIcRoSoFt.cOmPuTe/lOcAtIoNs/lOcAtIoNvAlUe/sHaReDgAlLeRiEs/gAlLeRyUnIqUeVaLuE/iMaGeS/gAlLeRyImAgEvAlUe/vErSiOnS/gAlLeRyImAgEvErSiOnVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseVersionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}

		if actual.Location != v.Expected.Location {
			t.Fatalf("Expected %q but got %q for Location", v.Expected.Location, actual.Location)
		}

		if actual.GalleryUniqueName != v.Expected.GalleryUniqueName {
			t.Fatalf("Expected %q but got %q for GalleryUniqueName", v.Expected.GalleryUniqueName, actual.GalleryUniqueName)
		}

		if actual.GalleryImageName != v.Expected.GalleryImageName {
			t.Fatalf("Expected %q but got %q for GalleryImageName", v.Expected.GalleryImageName, actual.GalleryImageName)
		}

		if actual.GalleryImageVersionName != v.Expected.GalleryImageVersionName {
			t.Fatalf("Expected %q but got %q for GalleryImageVersionName", v.Expected.GalleryImageVersionName, actual.GalleryImageVersionName)
		}

	}
}

func TestSegmentsForVersionId(t *testing.T) {
	segments := VersionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("VersionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
