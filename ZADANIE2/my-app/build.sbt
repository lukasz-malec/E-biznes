val ScalatraVersion = "2.8.4"

ThisBuild / scalaVersion := "2.12.18"

libraryDependencies ++= Seq(
  "org.scalatra" %% "scalatra" % ScalatraVersion,
  "org.eclipse.jetty" % "jetty-server" % "9.4.54.v20240208",
  "org.eclipse.jetty" % "jetty-webapp" % "9.4.54.v20240208",
  "org.json4s" %% "json4s-jackson" % "4.0.7",
  "ch.qos.logback" % "logback-classic" % "1.5.19",
  "javax.servlet" % "javax.servlet-api" % "3.1.0" % "provided"
)