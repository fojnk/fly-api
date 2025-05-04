CREATE OR REPLACE FUNCTION public.hashpoint(point) RETURNS integer
   LANGUAGE sql IMMUTABLE
   AS 'SELECT hashfloat8($1[0]) # hashfloat8($1[1])';

CREATE OPERATOR CLASS public.point_hash_ops DEFAULT FOR TYPE point USING hash AS
   OPERATOR 1 ~=(point,point),
   FUNCTION 1 public.hashpoint(point);

CREATE TABLE pricing_rules (
  flight_id INTEGER NOT NULL,
  fare_conditions VARCHAR(10) NOT NULL,
  min_price NUMERIC(10, 2) NOT NULL,
  max_price NUMERIC(10, 2) NOT NULL,
  avg_price NUMERIC(10, 2) NOT NULL,
  rule_description TEXT
);

INSERT INTO pricing_rules (flight_id, fare_conditions, min_price, max_price, avg_price)
SELECT
  f.flight_id,
  tf.fare_conditions,
  MIN(tf.amount) AS min_price,
  MAX(tf.amount) AS max_price,
  AVG(tf.amount) AS avg_price,
FROM
  flights f
JOIN
  ticket_flights tf ON f.flight_id = tf.flight_id
WHERE
  f.status <> 'Scheduled'
GROUP BY
  f.flight_id, tf.fare_conditions;