Fetch by batch 
‘’’Java
@Repository
public interface PostRepository extends Base]paRepository<Post, Long> {
@Query("'" select p from Post p
where date(p.createdon) >= :sinceDate
@QueryHints(
@QueryHint (name = AvailableHints. HINT_FETCH_SIZE, value = "25")
Stream‹Post> streambyCreatedonSince(@Param("sinceDate") LocalDate sinceDate);
} The FETCH_SIZE JPA query hint is necessary for PostgreSQL and MySQL to instruct the JDBC Driver to prefetch at most 25 records. Otherwise, the PostgreSQL and MySQL JDBC Drivers would prefetch all the query results
prior to traversing the underlying Resultset.
‘’’

Link - https://www.linkedin.com/safety/go?url=https%3A%2F%2Flnkd.in%2FdySaGfwy%3Ftrk%3Dfeed_main-feed-card-text

#spring #jpa #stream #database 
#draft