# Changelog

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.5.2] - 2025-01-27 Mon

Fis: taxon.ParentID does not show.

## [v0.5.1] - 2025-01-23 Thu

Add: update modules.
Add: document describing DwC terms used by TaxonWorks.
Fix: DwC terms.

## [v0.5.0] - 2025-01-20 Mon

Add [#12]: compatibility with CoLDP-based SFGA schema.
Add: cardinality data.
Add: option to process, skip or break on rows with wrong number of fields
Add: localID and globalID

## [v0.0.4] - 2024-09-09 Mon

Add: reset SFGA versions to v0.x.x (current v0.2.6) to allow backwards
incompatibility according to Semantic Versioning approach.
Add [#11]: migrate SQLite functionality to sflib.

## [v0.0.3] - 2024-08-21 Wed

Add [#10]: sfga is now in sflib.

## [v0.0.2] - 2024-04-04 Thu

Add: update dwca to 0.2.6 to skip normalization for already normalized files.

## [v0.0.1] - 2024-03-15 Fri

Add [#9]: ability to download DwCA from URL.
Add [#8]: export sfga to a dump file.
Add [#7]: ingest data_source data to sfga.
Add [#6]: ingest vernacular names to sfga.
Add [#5]: ingest Core to sfga.
Add [#4]: ingest DwCA file.
Add [#3]: import schema to sqlite work database.
Add [#2]: fetsch corresponding SFGA schema.
Add [#1]: create empty interfaces.

## [v0.0.0] - 2024-03-08 Fri

Add: initial commit

## Footnotes

This document follows [changelog guidelines]

[v0.5.1]: https://github.com/sfborg/from-dwca/compare/v0.5.0...v0.5.1
[v0.5.0]: https://github.com/sfborg/from-dwca/compare/v0.0.4...v0.5.0
[v0.0.4]: https://github.com/sfborg/from-dwca/compare/v0.0.3...v0.0.4
[v0.0.3]: https://github.com/sfborg/from-dwca/compare/v0.0.2...v0.0.3
[v0.0.2]: https://github.com/sfborg/from-dwca/compare/v0.0.1...v0.0.2
[v0.0.1]: https://github.com/sfborg/from-dwca/compare/v0.0.0...v0.0.1
[v0.0.0]: https://github.com/sfborg/from-dwca/tree/v0.0.0
[#20]: https://github.com/sfborg/from-dwca/issues/20
[#19]: https://github.com/sfborg/from-dwca/issues/19
[#18]: https://github.com/sfborg/from-dwca/issues/18
[#17]: https://github.com/sfborg/from-dwca/issues/17
[#16]: https://github.com/sfborg/from-dwca/issues/16
[#15]: https://github.com/sfborg/from-dwca/issues/15
[#14]: https://github.com/sfborg/from-dwca/issues/14
[#13]: https://github.com/sfborg/from-dwca/issues/13
[#12]: https://github.com/sfborg/from-dwca/issues/12
[#11]: https://github.com/sfborg/from-dwca/issues/11
[#10]: https://github.com/sfborg/from-dwca/issues/10
[#9]: https://github.com/sfborg/from-dwca/issues/9
[#8]: https://github.com/sfborg/from-dwca/issues/8
[#7]: https://github.com/sfborg/from-dwca/issues/7
[#6]: https://github.com/sfborg/from-dwca/issues/6
[#5]: https://github.com/sfborg/from-dwca/issues/5
[#4]: https://github.com/sfborg/from-dwca/issues/4
[#3]: https://github.com/sfborg/from-dwca/issues/3
[#2]: https://github.com/sfborg/from-dwca/issues/2
[#1]: https://github.com/sfborg/from-dwca/issues/1
[changelog guidelines]: https://keepachangelog.com/en/1.0.0/
