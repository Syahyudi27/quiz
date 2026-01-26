-- 1
CREATE SCHEMA IF NOT EXISTS oe;


--2
SELECT 
    c.category_id, 
    c.category_name, 
    COUNT(p.product_id) AS total_product
FROM 
    oe.categories c
LEFT JOIN 
    oe.products p ON c.category_id = p.category_id
GROUP BY 
    c.category_id, 
    c.category_name
ORDER BY 
    c.category_id;

--3
SELECT 
    s.supplier_id, 
    s.company_name, 
    COUNT(p.product_id) AS total_product
FROM 
    oe.suppliers s
LEFT JOIN 
    oe.products p ON s.supplier_id = p.supplier_id
GROUP BY 
    s.supplier_id, s.company_name
ORDER BY 
    total_product DESC;

--4
SELECT 
    s.supplier_id, 
    s.company_name, 
    COUNT(p.product_id) AS total_product,
    -- Menggunakan to_char untuk memformat rata-rata harga dengan 2 angka di belakang koma
    to_char(AVG(p.unit_price), 'FM999,999.00') AS avg_unit_price
FROM 
    oe.suppliers s
JOIN 
    oe.products p ON s.supplier_id = p.supplier_id
GROUP BY 
    s.supplier_id, 
    s.company_name
ORDER BY 
    total_product DESC;


--5
SELECT 
    p.product_id, 
    p.product_name, 
    p.supplier_id, 
    s.company_name, 
    p.unit_price, 
    p.units_in_stock, 
    p.units_on_order, 
    p.reorder_level
FROM 
    oe.products p
JOIN 
    oe.suppliers s ON p.supplier_id = s.supplier_id
WHERE 
    p.units_in_stock <= p.reorder_level
ORDER BY 
    p.product_name ASC;

--6
SELECT 
    c.customer_id, 
    c.company_name, 
    COUNT(o.order_id) AS total_order
FROM 
    oe.customers c
LEFT JOIN 
    oe.orders o ON c.customer_id = o.customer_id
GROUP BY 
    c.customer_id, 
    c.company_name
ORDER BY 
    c.customer_id ASC;

--7
SELECT 
    order_id, 
    customer_id, 
    order_date, 
    required_date, 
    shipped_date,
    -- Menghitung selisih hari sebagai delivery_time
    (shipped_date - order_date) AS delivery_time
FROM 
    oe.orders
WHERE 
    shipped_date IS NOT NULL 
    AND (shipped_date - order_date) > 7
ORDER BY 
    order_id ASC;

--8
SELECT 
    p.product_id, 
    p.product_name, 
    SUM(od.quantity) AS total_qty
FROM 
    oe.products p
JOIN 
    oe.order_details od ON p.product_id = od.product_id
GROUP BY 
    p.product_id, 
    p.product_name
ORDER BY
	total_qty DESC;

--9
SELECT 
    c.category_id, 
    c.category_name, 
    SUM(od.quantity) AS total_qty_ordered
FROM 
    oe.categories c
JOIN 
    oe.products p ON c.category_id = p.category_id
JOIN 
    oe.order_details od ON p.product_id = od.product_id
GROUP BY 
    c.category_id, 
    c.category_name
ORDER BY 
    total_qty_ordered DESC;

--10
WITH CategorySales AS (
    SELECT 
        c.category_id, 
        c.category_name, 
        SUM(od.quantity) AS total_qty_ordered
    FROM oe.categories c
    JOIN oe.products p ON c.category_id = p.category_id
    JOIN oe.order_details od ON p.product_id = od.product_id
    GROUP BY c.category_id, c.category_name
)
SELECT * FROM CategorySales
WHERE total_qty_ordered = (SELECT MAX(total_qty_ordered) FROM CategorySales)
   OR total_qty_ordered = (SELECT MIN(total_qty_ordered) FROM CategorySales)
ORDER BY total_qty_ordered DESC;

--11
SELECT 
    s.shipper_id, 
    s.company_name, 
    p.product_id, 
    p.product_name, 
    SUM(od.quantity) AS total_qty_ordered
FROM 
    oe.shippers s
JOIN 
    oe.orders o ON s.shipper_id = o.ship_via
JOIN 
    oe.order_details od ON o.order_id = od.order_id
JOIN 
    oe.products p ON od.product_id = p.product_id
GROUP BY 
    s.shipper_id, 
    s.company_name, 
    p.product_id, 
    p.product_name
ORDER BY 
    s.shipper_id DESC;

--12
SELECT shipper_id, company_name, product_id, product_name, total_qty_ordered
FROM (
    SELECT s.shipper_id, s.company_name, p.product_id, p.product_name, SUM(od.quantity) AS total_qty_ordered
    FROM oe.shippers s
    JOIN oe.orders o ON s.shipper_id = o.ship_via
    JOIN oe.order_details od ON o.order_id = od.order_id
    JOIN oe.products p ON od.product_id = p.product_id
    GROUP BY s.shipper_id, s.company_name, p.product_id, p.product_name
) main
WHERE total_qty_ordered = (
    SELECT MAX(total) FROM (
        SELECT SUM(od_max.quantity) as total FROM oe.orders o_max 
        JOIN oe.order_details od_max ON o_max.order_id = od_max.order_id 
        WHERE o_max.ship_via = main.shipper_id GROUP BY od_max.product_id
    ) s1
)
OR total_qty_ordered = (
    SELECT MIN(total) FROM (
        SELECT SUM(od_min.quantity) as total FROM oe.orders o_min 
        JOIN oe.order_details od_min ON o_min.order_id = od_min.order_id 
        WHERE o_min.ship_via = main.shipper_id GROUP BY od_min.product_id
    ) s2
)
ORDER BY shipper_id, total_qty_ordered ASC;
