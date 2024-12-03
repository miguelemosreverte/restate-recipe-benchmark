import os
import re

# Path to the reports directory
reports_dir = "reports"

# Output index file
output_file = "index.html"

# Regex to match and clean up filenames
def transform_filename(filename):
    # Remove "benchmark-report" and date-time stamps
    cleaned = re.sub(r"benchmark-report-\d{4}-\d{2}-\d{2}-\d{2}-\d{2}-\d{2}-?", "", filename)
    # Replace dashes with spaces
    cleaned = cleaned.replace("-", " ")
    # Capitalize the first letter of each word
    cleaned = cleaned.strip().capitalize()
    # Support percentages with high precision
    cleaned = re.sub(r"near ([0-9\.]+)percent", r"near \1 percent", cleaned)
    # Return None if the name becomes empty after cleaning
    return cleaned if cleaned.strip() else None

# Get the list of reports
reports = [f for f in os.listdir(reports_dir) if f.endswith(".html")]

# Generate the index.html content
index_content = """
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Benchmark Reports</title>
    <script src="https://cdn.tailwindcss.com"></script>
</head>
<body class="bg-gray-50 text-gray-800">
    <div class="container mx-auto px-4 py-8">
        <h1 class="text-4xl font-bold mb-6 text-center text-blue-600">Benchmark Reports</h1>
        <p class="text-center mb-4 text-gray-600">Explore the detailed benchmark reports below:</p>
        <ul class="space-y-4">
"""

for report in reports:
    presentable_name = transform_filename(report)
    if presentable_name != ".html":  # Skip files that result in empty names
        index_content += f'            <li class="bg-white shadow-md rounded-lg hover:shadow-lg transition-shadow duration-300">\n'
        index_content += f'                <a href="{reports_dir}/{report}" class="block px-6 py-4 text-lg font-medium text-blue-500 hover:text-blue-700">{presentable_name}</a>\n'
        index_content += f'            </li>\n'

index_content += """
        </ul>
    </div>
</body>
</html>
"""

# Write to the index.html file
with open(output_file, "w") as file:
    file.write(index_content)

print(f"Index file with Tailwind styling generated: {output_file}")
