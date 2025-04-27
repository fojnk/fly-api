CREATE TABLE pricing_rules (
  flight_id INTEGER NOT NULL,
  fare_conditions VARCHAR(10) NOT NULL,
  min_price NUMERIC(10, 2) NOT NULL,
  max_price NUMERIC(10, 2) NOT NULL,
  avg_price NUMERIC(10, 2) NOT NULL,
  rule_description TEXT
);

INSERT INTO pricing_rules (flight_id, fare_conditions, min_price, max_price, avg_price, rule_description)
SELECT
  f.flight_id,
  tf.fare_conditions,
  MIN(tf.amount) AS min_price,
  MAX(tf.amount) AS max_price,
  AVG(tf.amount) AS avg_price,
  'Standard pricing rule based on historical data' AS rule_description
FROM
  flights f
JOIN
  ticket_flights tf ON f.flight_id = tf.flight_id
WHERE
  f.status = 'Scheduled'
GROUP BY
  f.flight_id, tf.fare_conditions;