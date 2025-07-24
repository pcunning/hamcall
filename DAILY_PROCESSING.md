# Daily File Processing

This document describes the enhanced daily file processing feature added to the hamcall project.

## Overview

The system now supports two modes of daily file processing from the FCC ULS database:
- **Partial Mode**: Downloads only the most recent daily delta files (previous business day)
- **Full Mode**: Downloads the complete weekly files plus all daily delta files since the last weekly publication

This can significantly reduce download time and processing overhead for incremental updates while providing flexibility for different use cases.

## Usage

### Partial Mode (Delta Only)
Downloads only the previous business day's delta files:
```bash
ULS_MODE=partial ./hamcall -dl
```

### Full Mode (Weekly + All Dailies)
Downloads the complete weekly files plus all daily files since the weekly publication:
```bash
ULS_MODE=full ./hamcall -dl
```

### Weekly Mode (Default)
If `ULS_MODE` is not set or set to any other value, the system uses the default weekly file processing:
```bash
./hamcall -dl
```

## How It Works

### Daily File Timing
- Daily files are published the day after the data day (Monday's file available Tuesday after noon Eastern)
- Weekend handling: Saturday gets Friday's file, Sunday gets Friday's file
- Only business day files are available (Monday-Friday data)

### File Sources
Daily files are downloaded via FTP from:
- License data: `ftp://wirelessftp.fcc.gov:21/pub/uls/daily/l_am_{day}.zip`
- Application data: `ftp://wirelessftp.fcc.gov:21/pub/uls/daily/a_am_{day}.zip`
- Where `{day}` is the 3-letter day code (mon, tue, wed, thu, fri)

### Weekly File Schedule
- Weekly files are published on Sunday
- Full mode calculates all business days since the last Sunday to download

### Processing Modes

#### Partial Mode Behavior
- **ULS Data**: Processes only the daily delta files
- **Other Data Sources**: Only updates existing callsigns, does not create new entries
- **Purpose**: Preserves existing data from previous full runs while applying incremental changes

#### Full Mode Behavior
- **ULS Data**: Merges weekly files with all daily files since the weekly publication
- **Other Data Sources**: Normal processing, can create new callsigns
- **Purpose**: Complete refresh with all recent changes

### Fallback Mechanism
In partial mode, if daily files fail to download, unzip, or don't contain essential data (`AM.dat`, `EN.dat`, `HD.dat`), the system automatically falls back to weekly files.

## Data Source Behavior

### ULS Processing
- Same processing logic for all modes
- Full mode merges weekly + daily files before processing

### Other Data Sources (GEO, LOTW, RadioID)
- **Weekly/Full Mode**: Normal behavior, creates new callsigns if not in ULS data
- **Partial Mode**: Only updates existing callsigns, skips callsigns not in ULS data to preserve existing data

## Files Modified

- `main.go`: Updated to pass partial mode flag to data sources
- `source/uls/uls.go`: Enhanced with new modes, FTP downloads, and day calculation logic
- `source/geo/geo.go`: Added partial mode support
- `source/lotw/lotw.go`: Added partial mode support  
- `source/radioid/radioid.go`: Added partial mode support
- `source/uls/uls_test.go`: Updated tests for new functionality

## Functions Added

- `getPreviousBusinessDay()`: Calculates the most recent business day with available daily files
- `getAllDailysSinceWeekly()`: Gets all business days since the last weekly publication
- `mergeDailyFiles()`: Merges daily files with weekly files in full mode
- Updated `DownloadDailyLicenses()` and `DownloadDailyApplications()` for enhanced functionality

## Benefits

- **Flexible Processing**: Choose between quick delta updates or complete refreshes
- **Efficient Bandwidth**: Partial mode uses minimal bandwidth for regular updates
- **Complete Coverage**: Full mode ensures no changes are missed
- **Data Preservation**: Partial mode preserves existing data while applying incremental changes
- **Automatic Fallback**: Robust error handling with fallback to weekly files
- **Backward Compatibility**: Default behavior unchanged

## Testing

The implementation includes comprehensive tests:
- Unit tests for day calculation logic
- Tests for partial mode behavior
- Integration tests for mode selection
- Existing processing tests continue to pass