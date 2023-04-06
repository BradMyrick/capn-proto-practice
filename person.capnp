using Go = import "./std/go.capnp";
@0x961334376ded49ac;
$Go.package("main");
$Go.import("person");

struct Person {
  id @0 :UInt64;
  name @1 :Text;
  age @2 :UInt16;
  email @3 :Text;
}

struct AddressBook {
  people @0 :List(Person);
}
