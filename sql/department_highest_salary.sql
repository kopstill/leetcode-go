/*
    184: https://leetcode.cn/problems/department-highest-salary/
 */

SELECT
    Department.name AS 'Department',
    Employee.name AS 'Employee',
    Salary
FROM
    Employee
        JOIN
    Department ON Employee.DepartmentId = Department.Id
WHERE
    (Employee.DepartmentId , Salary) IN
    (   SELECT
            DepartmentId, MAX(Salary)
        FROM
            Employee
        GROUP BY DepartmentId
	)
;

-- Rank() | Dense_Rank() | Row_Number() 函数， MySQL >= 8.0
SELECT S.NAME 'Department', S.EMPLOYEE, S.SALARY
  FROM (SELECT D.NAME,
               T.NAME EMPLOYEE,
               T.SALARY,
               RANK() OVER(PARTITION BY T.DEPARTMENTID ORDER BY T.SALARY DESC) R
          FROM EMPLOYEE T
          LEFT JOIN DEPARTMENT D
            ON T.DEPARTMENTID = D.ID) S
 WHERE S.R = 1;
