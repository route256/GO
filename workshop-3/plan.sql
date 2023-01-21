set max_parallel_workers_per_gather = 0;

select a.aircraft_code,
               (
                   select round(avg(tf.amount))
                   from bookings.flights f
                            join bookings.ticket_flights tf on tf.flight_id = f.flight_id
                   where f.aircraft_code = a.aircraft_code AND
                         tf.amount > (select min(amount) from bookings.ticket_flights) AND
                         tf.amount < (select max(amount) from bookings.ticket_flights)

               )
from bookings.aircrafts a
group by a.aircraft_code;

create index on bookings.ticket_flights(amount);

begin isolation level repeatable read;

select max(amount) as a_max, min(amount) as a_min from bookings.ticket_flights \gset

explain (analyze, timing off) select a.aircraft_code,
               (
                   select round(avg(tf.amount))
                   from bookings.flights f
                            join bookings.ticket_flights tf on tf.flight_id = f.flight_id
                   where f.aircraft_code = a.aircraft_code AND
                         tf.amount > :a_min AND
                         tf.amount < :a_max

               )
from bookings.aircrafts a
group by a.aircraft_code;

explain (analyze, timing off) select a.aircraft_code, round(avg(tf.amount))
from bookings.aircrafts a
    left join bookings.flights f on f.aircraft_code = a.aircraft_code
    left join bookings.ticket_flights tf on tf.flight_id = f.flight_id AND
                                            tf.amount > :a_min AND
                                            tf.amount < :a_max
group by a.aircraft_code;

set work_mem = '8MB';

explain (analyze, costs off, timing off, buffers) select a.aircraft_code, round(avg(tf.amount))
from bookings.aircrafts a
    left join bookings.flights f on f.aircraft_code = a.aircraft_code
    left join bookings.ticket_flights tf on tf.flight_id = f.flight_id AND
                                            tf.amount > :a_min AND
                                            tf.amount < :a_max
group by a.aircraft_code;