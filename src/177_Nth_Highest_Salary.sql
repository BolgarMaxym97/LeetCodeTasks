Create table If Not Exists Employee (Id int, Salary int);
Truncate table Employee;
insert into Employee (id, salary) values ('1', '100');
insert into Employee (id, salary) values ('2', '200');
insert into Employee (id, salary) values ('3', '300');



CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
RETURN (
    with cte as
             (select salary
                   , ROW_NUMBER() over (order by salary desc ) as rn
              from Employee
              group by salary)

    select IFNULL((select salary
                   from cte
                   where rn = N), null) as SecondHighestSalary
    );
END;

getNthHighestSalary(2);


