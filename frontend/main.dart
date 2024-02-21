import 'dart:io';

void main() async {
  // TODO: hardcoded
  print("Connecting to TCP server...");
  final sock = await Socket.connect("127.0.0.1", 8000);
  
  sock.listen((bytes) {
    print("Recieved data");
  });

  await sock.flush();
}