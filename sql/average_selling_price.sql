/*
 * 1251: https://leetcode.cn/problems/average-selling-price/
 */

-- Write your MySQL query statement below
select 
  p.product_id, 
  round(sum(p.price * us.units) / sum(us.units), 2) average_price 
from 
  Prices p, 
  UnitsSold us 
where 
  p.product_id = us.product_id and 
  us.purchase_date between p.start_date and p.end_date 
group by 
  p.product_id

-- subquery
SELECT
    product_id,
    Round(SUM(sales) / SUM(units), 2) AS average_price
FROM (
    SELECT
        Prices.product_id AS product_id,
        Prices.price * UnitsSold.units AS sales,
        UnitsSold.units AS units
    FROM Prices 
    JOIN UnitsSold ON Prices.product_id = UnitsSold.product_id
    WHERE UnitsSold.purchase_date BETWEEN Prices.start_date AND Prices.end_date
) T
GROUP BY product_id
