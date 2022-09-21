/*
    177: https://leetcode.cn/problems/nth-highest-salary/
 */

-- set global log_bin_trust_function_creators = 1;

-- show variables like 'log_bin_trust_function_creators';

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    SET N := N - 1;
  RETURN (
      -- Write your MySQL query statement below.
      SELECT
            salary
      FROM
            employee
      GROUP BY
            salary
      ORDER BY
            salary DESC
      LIMIT N, 1
  );
END

-- select getNthHighestSalary(3);

-- drop function getNthHighestSalary;

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    SET N := N - 1;
  RETURN (
      -- Write your MySQL query statement below.
      SELECT DISTINCT
            salary
      FROM
            employee
      ORDER BY
            salary DESC
      LIMIT N, 1
  );
END
