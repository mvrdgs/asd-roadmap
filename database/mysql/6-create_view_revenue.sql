CREATE VIEW current_revenue AS
SELECT
    (SELECT FORMAT(MIN(plan_price), 2) FROM plans) AS minimun_revenue,
    (SELECT FORMAT(MAX(plan_price), 2) FROM plans) AS maximun_revenue,
    FORMAT(
            (
                SELECT CEIL(SUM(plan_price) / (SELECT COUNT(*) FROM users) * 100) / 100 FROM plans
            ), 2
        ) AS faturamento_medio,
    FORMAT((SELECT SUM(plan_price) FROM plans), 2) AS total_revenue
FROM user_status AS us
GROUP BY total_revenue;
