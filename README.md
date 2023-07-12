# CapCut Scraper

## A simple web scraper especifically for CapCut

CapCut is an app that has pages dedicated to CapCut Templates. There are two types of pages
(template pages and template detail pages):

1. Example Template ID: 7156421419591634182
   a. https://www.capcut.com/templates/7156421419591634182
   b. https://www.capcut.com/template-detail/7156421419591634182
2. Example Template ID: 7194527218402446619
   a. https://www.capcut.com/template/7194527218402446619
   b. https://www.capcut.com/template-detail/7194527218402446619

Each of these templates is denoted by an ID in the URL which is the template ID.
The goal is to create an API that takes in one of those IDs and returns the scraped data from
the corresponding CapCut page. The API definition is listed below.
The scraped content of the page is in the form of a JSON response. The content of one
of the CapCut pages have a script tag with id RENDER_DATA and a type of
application/json `(<script id="RENDER_DATA" type="application/json">)`. It returns the
JSON that is embedded in that tag. The content of the second type of page is having a script
tag with content that starts with `<script>window._ROUTER_DATA =...</script>`. It
return the JSON that is embedded in that script tag.

### How to run

- Run server:

  ```bash
  make scrape
  ```

- Get GET /capcut/template/<template_id> :

  ```bash
    localhost:8080/capcut/template/7156421419591634182
  ```
- Get GET /capcut/template-detail/<template_id> :

  ```bash
    localhost:8080/capcut/template-detail/7156421419591634182
  ```
