import 'dart:io';
import 'dart:convert' show utf8;

import 'dart:typed_data';
import 'dart:core';

Uint8List numPack(int packNum) {
  final list = Uint8List(2);
  list[1] = packNum % 256;
  list[0] = (packNum / 256).floor();
  
  return list;
}

int numUnpack(Uint8List packedData) {
  return (packedData[0] * 256) + packedData[1];
}

void main() async {
  // TODO: hardcoded
  print("Connecting to TCP server...");
  final sock = await Socket.connect("127.0.0.1", 8000);

  sock.listen((bytes) {
    print(utf8.decode(bytes));
  });

  await sock.flush();
}
