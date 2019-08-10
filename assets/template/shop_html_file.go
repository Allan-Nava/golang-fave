package template

var VarShopHtmlFile = []byte(`{{template "header.html" .}}
<div class="card mb-4">
	{{if $.Data.Shop.HaveProducts}}
		{{range $.Data.Shop.Products}}
			<div class="post">
				<div class="card-body">
					<h2 class="card-title">
						<a href="{{.Permalink}}">
							{{.Name}}
						</a>
					</h2>
					<div class="post-content">
						{{.Briefly}}
					</div>
					<div class="post-date">
						<div><small>Price: {{.PriceFormat "%.2f"}} {{.Currency.Code}}</small></div>
						<div><small>Published on {{.DateTimeFormat "02/01/2006, 15:04:05"}}</small></div>
						<div>Author: {{.User.FirstName}} {{.User.LastName}}</div>
					</div>
				</div>
			</div>
		{{end}}
	{{else}}
		<div class="card-body">
			Sorry, no products matched your criteria
		</div>
	{{end}}
</div>
{{if $.Data.Shop.HaveProducts}}
	{{if gt $.Data.Shop.ProductsMaxPage 1 }}
		<nav>
			<ul class="pagination mb-4">
				{{if $.Data.Shop.PaginationPrev}}
					<li class="page-item{{if $.Data.Shop.PaginationPrev.Current}} disabled{{end}}">
						<a class="page-link" href="{{$.Data.Shop.PaginationPrev.Link}}">Previous</a>
					</li>
				{{end}}
				{{range $.Data.Shop.Pagination}}
					{{if .Dots}}
						<li class="page-item disabled"><a class="page-link" href="">...</a></li>
					{{else}}
						<li class="page-item{{if .Current}} active{{end}}">
							<a class="page-link" href="{{.Link}}">{{.Num}}</a>
						</li>
					{{end}}
				{{end}}
				{{if $.Data.Shop.PaginationNext}}
					<li class="page-item{{if $.Data.Shop.PaginationNext.Current}} disabled{{end}}">
						<a class="page-link" href="{{$.Data.Shop.PaginationNext.Link}}">Next</a>
					</li>
				{{end}}
			</ul>
		</nav>
	{{end}}
{{end}}
{{template "footer.html" .}}`)
