val ScalatraVersion = "2.8.4"

ThisBuild / scalaVersion := "2.13.12"

lazy val root = (project in file("."))
  .settings(
    name := "scalatra-app",
    version := "0.1.0",

    libraryDependencies ++= Seq(
      "org.scalatra" %% "scalatra" % ScalatraVersion,
      "ch.qos.logback" % "logback-classic" % "1.5.19",

      "org.eclipse.jetty" % "jetty-server" % "9.4.54.v20240208",
      "org.eclipse.jetty" % "jetty-webapp" % "9.4.54.v20240208",

      "org.json4s" %% "json4s-jackson" % "4.0.7",
      "ch.qos.logback" % "logback-classic" % "1.5.19"
      // "org.scalatra" %% "scalatra-json-jackson" % "2.8.4"
    )
  )