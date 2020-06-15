import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Food Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        visualDensity: VisualDensity.adaptivePlatformDensity,
      ),
      home: MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        decoration: BoxDecoration(
          gradient: LinearGradient(colors: [Colors.pink, Colors.blue]),
        ),
        child: Container(
          width: MediaQuery.of(context).size.width,
          height: MediaQuery.of(context).size.height,
          margin: EdgeInsets.symmetric(horizontal: 50),
          alignment: Alignment.center,
          child: Center(
            child: Card(
              margin: EdgeInsets.symmetric(vertical: 85.0),
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.all(Radius.circular(25)),
              ),
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: <Widget>[
                  Text(
                    'SIGN IN',
                    style: TextStyle(fontSize: 30.0),
                  ),
                  const SizedBox(height: 30),
                  TextField(
                    decoration: InputDecoration(
                      fillColor: Colors.pink,
                      border: new OutlineInputBorder(
                        borderRadius: const BorderRadius.all(
                          const Radius.circular(25.0),
                        ),
                      ),
                      prefixIcon: Icon(
                        Icons.email,
                        color: Colors.black,
                        size: 30.0,
                      ),
                    ),
                  ),
                  const SizedBox(height: 30),
                  TextField(
                    obscureText: true,
                    decoration: InputDecoration(
                      fillColor: Colors.pink,
                      border: OutlineInputBorder(
                        borderRadius: const BorderRadius.all(
                          const Radius.circular(25.0),
                        ),
                      ),
                      prefixIcon: Icon(
                        Icons.lock,
                        color: Colors.black,
                        size: 30.0,
                      ),
                    ),
                  ),
                  const SizedBox(height: 30),
                  RaisedButton(
                    onPressed: () {},
                    child: Text(
                      'Login'.toUpperCase(),
                      style: TextStyle(color: Colors.white),
                    ),
                    color: Colors.pink,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.all(Radius.circular(25)),
                    ),
                  )
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
