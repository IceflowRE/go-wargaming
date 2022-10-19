*********
Changelog
*********

All notable changes to this project will be documented in this file.

The format is based on `Keep a Changelog <https://keepachangelog.com/en/1.0.0/>`_ and this project adheres to `Semantic Versioning <https://semver.org/spec/v2.0.0.html>`_.

3.0.0
=====

Update
------

- Updated API

  Some API endpoints were removed.

Changed
-------

- Some struct fields got renamed.

  Example: ``Type_`` -> ``Type``

Removed
-------

- russian realm

2.2.3
=====

Updated
-------

- Updated API

2.2.2
=====

Fixed
-----

- nil pointer dereference in ``NewClient``

2.2.1
=====

Updated
-------

- Improved documentation

2.2.0
=====

Changed
-------

- made services public

2.1.0
=====

Fix
---

- go module import path

2.0.0
=====

Replaced
--------

- ``ApiErrorStringToString(error) string`` with ``ResponseError.Description() string``

Changed
-------

- ``wot/globalmap/eventaccountinfo`` made ``account_id`` optional (either ``account_id`` or ``clan_id`` is required.
