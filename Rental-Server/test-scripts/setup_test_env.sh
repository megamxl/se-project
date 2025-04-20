#!/bin/bash

echo "🔧 Setting up Python test environment..."

# Step 1: Create venv if not exists
if [ ! -d "venv" ]; then
    echo "📦 Creating virtual environment..."
    python3 -m venv venv
else
    echo "📁 Virtual environment already exists."
fi

# Step 2: Activate the venv
echo "✅ Activating virtual environment..."
source venv/bin/activate

# Step 3: Install dependencies
echo "⬇️ Installing Python dependencies..."
pip install --upgrade pip
pip install requests

echo "🎉 Environment is ready! You can now run:"
echo "   source venv/bin/activate"
echo "   python test_cars.py"