# TaxonWorks DwCA

This document describes DwCA and TW terms needed for TaxonWorks DwCA import

## Terms links

[Data Mapping](https://docs.taxonworks.org/guide/import.html#map-your-data)

Data is collected to `DwcOccurrence` table

```json
{
  "id": 7835569,
  "basisOfRecord": "PreservedSpecimen",
  "catalogNumber": "INHS Insect Collection 3792",
  "dwcClass": "Insecta",
  "coordinateUncertaintyInMeters": "10.0",
  "country": "United States",
  "county": "Shelby",
  "dateIdentified": "2002",
  "day": "17",
  "decimalLatitude": "39.232285",
  "decimalLongitude": "-88.82468",
  "eventDate": "2002-02-17",
  "family": "Capniidae",
  "footprintWKT": "POINT Z (-88.82468 39.232285 0)",
  "genus": "Allocapnia",
  "georeferenceProtocol": "A geospatial point translated from verbatim values recorded on human-readable media (e.g. paper specimen label, field notebook).",
  "georeferenceRemarks": "Derived from a instance of TaxonWorks' Georeference::VerbatimData.",
  "georeferenceSources": "Physical collection object.",
  "georeferencedBy": "Ed Dewalt",
  "higherClassification": "Animalia | Arthropoda | Insecta | Plecoptera | Arctoperlaria | Nemouroidea | Capniidae | Allocapnia",
  "individualCount": 1,
  "institutionCode": "INHS",
  "institutionID": "http://grbio.org/institution/illinois-natural-history-survey",
  "kingdom": "Animalia",
  "lifeStage": "Adult(s)",
  "month": "2",
  "nomenclaturalCode": "iczn",
  "occurrenceID": "2b2794a2-d5e7-4091-8672-c3fa6850d53c",
  "occurrenceStatus": "present",
  "order": "Plecoptera",
  "phylum": "Arthropoda",
  "preparations": "Vial",
  "recordedBy": "J.E. Petzing, J.W. Petzing",
  "scientificName": "Allocapnia vivipara (Claassen, 1924)",
  "scientificNameAuthorship": "(Claassen, 1924)",
  "sex": "Female",
  "specificEpithet": "vivipara",
  "startDayOfYear": "48",
  "stateProvince": "Illinois",
  "taxonRank": "species",
  "verbatimCoordinates": "N39.232285º W88.82468º",
  "verbatimLatitude": "N39.232285º",
  "verbatimLocality": "1.6 km N Holliday",
  "verbatimLongitude": "W88.82468º",
  "waterBody": "Richland Creek",
  "year": "2002",
  "dwc_occurrence_object_id": 210877,
  "dwc_occurrence_object_type": "CollectionObject",
  "project_id": 1,
  "created_at": "2024-09-10T14:11:45.393Z",
  "updated_at": "2025-01-18T00:53:45.644Z",
  "verbatimLabel": "IL: Shelby Co.; Richland; Cr. @ Co.-Rd. 1700E, 1.6 km; N Holliday. Z16 342503 E; 4343918 N; 17 Feb. 2002; J. W. Petzing & J. E. Petzing\n\nAllocapnia vivipara, 1 F; Det. R. E. DeWalt, Feb. 2002",
  "superfamily": "Nemouroidea"
}
```

All terms are DwC except

```ruby
TW_ATTRIBUTES = [
    :id,
    :project_id,
    :created_at,
    :updated_at,
    :created_by_id,
    :updated_by_id,
    :dwc_occurrence_object_type,
    :dwc_occurrence_object_id
  ].freeze
```

where `dwc_occurrence_object_type` can be one of
`CollectionObject, AssertedDistribution, FieldOccurrence`

## Classification Checklist

[GBIF relevant doc](https://ipt.gbif.org/manual/en/ipt/latest/best-practices-checklists)

```text
acceptedNameUsageID
namePublishedInYear
nomenclaturalCode
originalNameUsageID
parentNameUsageID
scientificName
scientificNameAuthorship
taxonID
taxonomicStatus
taxonRank
TW:TaxonNameClassification:Iczn:Fossil
TW:TaxonNameClassification:Latinized:Gender
TW:TaxonNameClassification:Latinized:PartOfSpeech
TW:TaxonNameRelationship:incertae_sedis_in_rank
```
