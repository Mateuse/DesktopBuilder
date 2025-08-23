-- Convert all brand names to lowercase in the components table
-- This ensures consistency with the new scraper behavior

-- First, let's see what brands we have and how many components each has
-- (Uncomment the following lines to see the current state)
-- SELECT brand, COUNT(*) as count 
-- FROM components 
-- GROUP BY brand 
-- ORDER BY brand;

-- Update all brands to lowercase
UPDATE components 
SET brand = LOWER(brand);

-- Verify the update
SELECT 'Conversion completed. Updated brands:' as message;
SELECT DISTINCT brand, COUNT(*) as count 
FROM components 
GROUP BY brand 
ORDER BY brand;
