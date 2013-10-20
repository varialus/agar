Agar
====

Translation of the Agar Widget Toolkit from C to Pure Go

Translation Ideas
-----------------

* Curent Plan: Use cgo initially, then translate from C to Go.
* Fix errors as they occur by translating from C to Go.
* Translate everything all at once from C to Go, then fix errors.

Issues
------

* When translating straight from C to Go it feels inconsistent to make case by case judgment calls whether to convert pointers to values.
