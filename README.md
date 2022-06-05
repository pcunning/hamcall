# hamcall.dev

This project creates a service to lookup of amateur radio call signs that is fast, scalable and simple.

We do this by generating a file for each amateur radio call sign that contains the information about the call sign from the FCC ULS database and other interesting sources. These files are uploaded to [Backblaze B2](https://www.backblaze.com/b2) (a S3 like service) and then served through [Cloudflare](https://www.cloudflare.com) as a CDN. Cloudflare provides URL rewriting between the public URL format and Backblaze public URL format.

Because of the project architecture, it is extremely inexpensive and very scalable. Please feel free to use hamcall.dev for any purpose. If you have any questions, notice issues with the data, or have an idea for improvement please open an issues.

## Why B2, Cloudflare and GitHub Actions

Its inexpensive, hosting all files for this site in [B2 costs](https://www.backblaze.com/b2/cloud-storage-pricing.html) about $0.06 per year (Backblaze also gives away the first 10GB for free). When placed behind Cloudflare's (a member of the Bandwidth Alliance) [generous free tier](https://www.cloudflare.com/plans/#add-ons) egress is free. Lastly GitHub Actions is provided [free for public projects](https://docs.github.com/en/billing/managing-billing-for-github-actions/about-billing-for-github-actions) allowing the build and upload of the files to happen [automatically for free on a schedule](https://github.com/pcunning/hamcall/blob/main/.github/workflows/run-build.yml#L6).

### B2 Configuration

The B2 bucket is configured to allow public read access and have permissive CORES headers.

### Cloudflare Configuration

While I initially thought I would need to use workers to rewite the B2 URLS to the format I wanted I found that I could do this with Cloudflare's [Tranform Rules](https://developers.cloudflare.com/rules/transform). 

I have three rules:

1. The SPA (index.html) is served from the root of the domain and any URL without a .suffix so the spa can load and subsequently load the correct callsign json file. 
    - `(not http.request.uri contains ".")` rewrite to static `file/hamcall/index.html`
2. Static assets to allow loading of images, css, js, etc. in the index.html.
    - `(http.request.uri contains "." and not http.request.uri contains ".json")` rewrite to dynamic `concat("/file/hamcall", lower(http.request.uri.path))`
3. JSON files (the API) that are separated to a separate path in B2 (makes the b2 bucket a little cleaner)
    - `(http.request.uri contains ".json")` rewrite to dynamic `concat("/file/hamcall/callsigns", lower(http.request.uri.path))`


## Process to get here

While I have looked at several other APIs to get similar data I couldn't find one that was simple and contained all the data I needed about a call sign (or more accurately a ULS license). While I could have put the data in a database and queried it from there it seemed like a good challenge to see how minimalist on the hosting side I could make this. For static generated sites I often reach for Netlify so I wrote a quick program to generate a bunch of json files and sent it off to Netlify hoping for the best. 

My initial build process wrote to a file in several passes for each ULS file it read. This took a lot of time and exceeded Netlify's 20 minute build time (IO is expensive, see [Latency Numbers Every Programer Should Know](https://colin-scott.github.io/personal_website/research/interactive_latency.html)). Since memory is cheap I ditched the IO and saved writing files until the end. How would Netlify handel it? While I could build all 1.5 Million files in about 5 min Netlify's CDN uploader would just crash trying to read the output directory. 

To be fair this is significantly outside of a normal use case for Netlify. Looking at other static site hosts like Cloudflare I would be limited to [20k files](https://developers.cloudflare.com/pages/platform/limits#files). So against better judgment I decided to gave it a try locally and 5 min later I had 1.5 million files in a directory. I spent the next 3 hours figuring out how to delete the files (rsync from an empty folder finally worked). Lets not do that again.

Since the hangup is moving the files to the CDN the next step was to write them directly to a CDN. Once I updated the program my first run took several hours to build and upload all 1.5 Million files (over residential internet) but thankfully I thought ahead and saved the hash for each file so subsequent runs can compare and only update changed files.

## Data Sources

- FCC ULS Database https://www.fcc.gov/uls/transactions/daily-weekly
- LOTW https://lotw.arrl.org/
- RadioID https://www.radioid.net/
- Location https://haminfo.tetranz.com