using Go = import "../std/go.capnp";
@0xdd30180cd0e36f1f;
$Go.package("receipt");
$Go.import("receipt");

struct Receipt {
  id @0 :UInt64;
  data @1 :Data;
  signature @2 :Data;
}

