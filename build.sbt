lazy val root = (project in file(".")).
  settings(
    inThisBuild(List(
      organization := "com.github.tanishiking",
      scalaVersion := "2.13.6",
      version      := "0.0.1"
    )),
    name := "ecs-kubejob-test"
  )

libraryDependencies ++= Seq(
  "org.scalatest"              %% "scalatest" % "3.2.9" % Test,
  "com.typesafe.scala-logging" %% "scala-logging"       % "3.9.4",
  "ch.qos.logback"              % "logback-classic"     % "1.2.7",
  "org.slf4j"                   % "slf4j-api"           % "1.7.32",
  "co.elastic.logging"          % "logback-ecs-encoder" % "1.2.0"
)

enablePlugins(JavaAppPackaging, DockerPlugin)

dockerBaseImage := "openjdk:11"
dockerUsername := Some("tanishiking")

