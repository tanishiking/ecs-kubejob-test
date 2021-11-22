import com.typesafe.scalalogging.LazyLogging

object Hello extends App with LazyLogging {
  val e = new Exception("foo")
  logger.info("foo") 
  logger.error(e.getMessage(), e)

}
