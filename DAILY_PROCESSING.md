# Daily File Processing

This document describes the daily file processing feature added to the hamcall project.

## Overview

The system now supports processing daily delta files from the FCC ULS database instead of the full weekly files. This can significantly reduce download time and processing overhead for incremental updates.

## Usage

Set the environment variable `ULS_USE_DAILY=true` to enable daily file processing:

```bash
ULS_USE_DAILY=true ./hamcall -dl
```

If the environment variable is not set or set to any value other than "true", the system will use the default weekly file processing.

## How It Works

1. **Daily File Download**: When `ULS_USE_DAILY=true`, the system attempts to download daily files:
   - License data: `https://data.fcc.gov/download/pub/uls/daily/l_am_{day}.zip` 
   - Application data: `https://data.fcc.gov/download/pub/uls/daily/a_am_{day}.zip`
   - Where `{day}` is the current day of week (sun, mon, tue, wed, thu, fri, sat)

2. **Fallback Mechanism**: If daily files fail to download, unzip, or don't contain essential data, the system automatically falls back to weekly files.

3. **Essential Data Check**: For license files, the system verifies that `AM.dat`, `EN.dat`, and `HD.dat` files exist and have content before proceeding.

## Files Modified

- `source/uls/uls.go`: Added daily download functions and logic
- `source/uls/uls_test.go`: Added tests for daily processing functionality

## Functions Added

- `DownloadDailyLicenses()`: Downloads daily license files with fallback
- `DownloadDailyApplications()`: Downloads daily application files with fallback  
- `dailyFilesContainEssentialData()`: Validates that essential files exist and have content

## Benefits

- **Faster Updates**: Daily files contain only changes since the last weekly update
- **Reduced Bandwidth**: Smaller file sizes for regular updates
- **Automatic Fallback**: Seamless fallback to weekly files if daily files are unavailable
- **Backward Compatibility**: Default behavior unchanged, weekly processing still the default

## Testing

The implementation includes comprehensive tests:
- Unit tests for daily file validation
- Integration tests for the download selection logic
- Existing processing tests continue to pass