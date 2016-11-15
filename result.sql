use poll_1
SELECT vote_0, COUNT(*) AS freq FROM ballot GROUP BY vote_0;