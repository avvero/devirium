Log slow query with threshold by adding this line to the configuration file(It available in Hibernate 5.4.5+), otherwise use third-party libraries such as _datasource-proxy._

hibernate.session.events.log.LOG_QUERIES_SLOWER_THAN_MS=30

#sql #hibernate #metric