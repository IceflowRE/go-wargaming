*********
Changelog
*********

All notable changes to this project will be documented in this file.

The format is based on `Keep a Changelog <https://keepachangelog.com/en/1.0.0/>`_ and this project adheres to `Semantic Versioning <https://semver.org/spec/v2.0.0.html>`_.

4.1.6
=====

Updated
-------

- wotx: update endpoint ``https://api-console.worldoftanks.??/wotx`` to ``https://api-modernarmor.worldoftanks.%s/wotx``

4.1.5
=====

Updated
-------

- wot: ``acount/info``
- wot: ``tank/stats``

Removed
-------

- wgn: ``wgtv/tags``
- wgn: ``wgtv/vehicles``
- wgn: ``wgtv/videos``

4.1.4
=====

Updated
-------

- Improved client code quality

Removed
-------

- wotb: ``clanmessages/create``
- wotb:``clanmessages/delete``
- wotb:``clanmessages/like``
- wotb:``clanmessages/likes``
- wotb:``clanmessages/messages``
- wotb:``clanmessages/update``

4.1.3
=====

Updated
-------

- Fix formatting of generated code (#8)

4.1.2
=====

Updated
-------

- Updated to Go 1.21

Changed
-------

- Generator has no underscore in front anymore.

Fixed
-----

- wot: ``encyclopedia/vehicles`` patch some return types

4.1.1
=====

Fixed
-----

- wot: ``account/tanks`` return type (again)
- wot: ``encyclopedia/vehicles`` return type

4.1.0
=====

Updated
-------

- wows api to recent version

Fixed
-----

- wot: ``account/tanks`` return type

4.0.1
=====

Changed
-------

- Removed generator from module

4.0.0
=====

Added
-----

- Introduced a ``GenericMeta`` struct which will be returned for specific endpoints. If some endpoints were overlooked, they will be added in a minor version.
Thanks at `@kakwa <https://github.com/kakwa>`_ for pointing this out.

Fixed
-----

- Fixed a nil pointer dereference if a nil value was passed at client creation.

3.0.0
=====

Updated
-------

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

- wot: ``globalmap/eventaccountinfo`` made ``account_id`` optional (either ``account_id`` or ``clan_id`` is required.
