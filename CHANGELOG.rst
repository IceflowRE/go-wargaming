*********
Changelog
*********

All notable changes to this project will be documented in this file.

The format is based on `Keep a Changelog <https://keepachangelog.com/en/1.0.0/>`_ and this project adheres to `Semantic Versioning <https://semver.org/spec/v2.0.0.html>`_.

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
