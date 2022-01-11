<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .Q.Name }}</title>
    <link href="/static/bootstrap-4.6.1-dist/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/fontawesome-free-5.15.4-web/css/all.min.css" rel="stylesheet">
    <link href="/static/bootstrap-table-1.19.1/bootstrap-table.min.css" rel="stylesheet">
    <link href="/static/select2-4.0.13/css/select2.min.css" rel="stylesheet">
    <link href="/static/select2-4.0.13/css/select2-bootstrap4.min.css" rel="stylesheet">
    <meta content="width=device-width, initial-scale=1, shrink-to-fit=no" name="viewport">
    <style type="text/css">
        .show-after {
            display: none;
        }
    </style>
</head>
<body>
<div class="container">
    <div class="card ">
        <div class="card-body border-white">
            <h5 class="card-title text-center text-primary">{{ .Q.Name }}</h5>
            <h6 class="card-subtitle  text-muted text-center">
                我们会对所填写的内容做到绝对保密，不会对任何其他单位或者非相关工作员工泄漏以下信息。对自己承诺负责。</h6>
            <br/>
            <div class="card-text">
                {{if .Message}}
                    <div class="alert alert-success">{{.Message}}</div>
                {{else}}
                    <form class="needs-validation" enctype="multipart/form-data" novalidate method="post">
                        {{range .Q.Questions}}
                            <input hidden type="text" id="q_{{.ID}}_use" name="q_{{.ID}}_use">
                            {{ if eq .Type "text"}}
                                <div class="form-group {{if .ForOtherID}}show-after{{end}}" data-rel="q_{{.ForOtherID}}"
                                >
                                    <label for="q_{{.ID}}">
                                        {{if .Require }}
                                            <strong class="text-danger">*</strong>
                                        {{end}}
                                        {{.Content}}： </label>
                                    <input class="form-control time-tick" id="q_{{.ID}}" data-required="{{.Require}}"
                                           name="q_{{.ID}}" type="text"
                                           {{if and .Require (not .ForOtherID)}}required{{end}}>
                                </div>
                            {{else if eq .Type "radio"}}
                                <div class="form-group">
                                    <label for="q_{{.ID}}">
                                        {{if .Require}}
                                            <strong class="text-danger">*</strong>
                                        {{end}}
                                        {{.Content}}：
                                    </label>
                                    <div class="row row-cols-1 row-cols-sm-2 row-cols-md-4 ml-3">
                                        {{$qid := .ID }}
                                        {{$r := .Require }}
                                        {{range .Options}}
                                            <div class="form-check  col">
                                                <input class="form-check-input time-tick" id="o_{{$qid}}_{{.ID}}"
                                                       name="q_{{$qid}}" data-other="{{.Other}}"
                                                       {{if $r}}required{{end}}
                                                       type="radio" value="{{.ID}}">
                                                <label class="form-check-label" for="o_{{$qid}}_{{.ID}}">
                                                    {{.Content}}
                                                </label>
                                            </div>
                                        {{end}}
                                    </div>
                                </div>
                            {{else if eq .Type "check"}}
                                <div class="form-group ">
                                    <label for="q_{{.ID}}">
                                        {{if .Require}}
                                            <strong class="text-danger">*</strong>
                                        {{end}}
                                        {{.Content}}：
                                    </label>
                                    <div class="row row-cols-1 row-cols-sm-2 row-cols-md-4 ml-3">
                                        {{$qid := .ID }}
                                        {{$r := .Require }}
                                        {{range .Options}}
                                            <div class="form-check  col">
                                                <input class="form-check-input time-tick" id="o_{{$qid}}_{{.ID}}"
                                                       name="q_{{$qid}}" data-other="{{.Other}}"
                                                       type="checkbox" value="{{.ID}}">
                                                <label class="form-check-label" for="o_{{$qid}}_{{.ID}}">
                                                    {{.Content}}
                                                </label>
                                            </div>
                                        {{end}}
                                    </div>
                                </div>
                            {{else if eq .Type "select"}}
                                <div class="form-group">
                                    <label for="q_{{.ID}}">
                                        {{if .Require}}
                                            <strong class="text-danger">*</strong>
                                        {{end}}
                                        {{.Content}}：
                                    </label><br/>
                                    <select class="custom-select time-tick" id="q_{{.ID}}" name="q_{{.ID}}"
                                            {{if .Require}}required{{end}}>
                                        <option disabled selected value="">请选择</option>
                                        {{$qid := .ID }}
                                        {{range .Options}}
                                            <option value="{{.ID}}" data-other="{{.Other}}">{{.Content}}</option>
                                        {{end}}
                                    </select>
                                </div>
                            {{else}}
                                <p>类型错误{{.Type}}</p>
                            {{end}}

                        {{end}}

                        <div class="text-center">
                            <button class="btn btn-primary btn-block" type="submit">提交</button>
                        </div>
                    </form>
                {{end}}
            </div>
        </div>
    </div>
</div>

<script src="/static/jquery-3.6.0.min.js"></script>
<script src="/static/bootstrap-4.6.1-dist/js/bootstrap.bundle.min.js"></script>
<script>

	let st = new Date();
	let start = st.getTime()
	let lastName = ""

	function doShowAfter() {
		let showAfter = $(".show-after")
		let showAfterInput = showAfter.find("input")
		showAfter.each(function () {
			let currentShow = $(this)
			let rel = $(":input[name='" + $(this).data("rel") + "']")
			if (rel.is("SELECT")) {
				rel.on("input", function (e) {
					if ($(e.delegateTarget).find("option:selected").data("other")) {
						currentShow.show()
						if (showAfterInput.data("required")) {
							showAfterInput.attr("required", true);
						}
					} else {
						currentShow.hide()
						if (showAfterInput.data("required")) {
							showAfterInput.removeAttr("required");
						}
						currentShow.removeAttr("required");
					}
				})
			} else {
				rel.each(function () {

					if ($(this).data("other")) {
						$(this).on("input", function (e) {
							if ($(e.delegateTarget).is(':checked')) {
								currentShow.show()
							} else {
								currentShow.hide()
							}
						})
					}
				})

			}
		});
	}

	function doVolidate() {
		// Fetch all the forms we want to apply custom Bootstrap validation styles to
		var forms = document.getElementsByClassName('needs-validation');
		// Loop over them and prevent submission
		var validation = Array.prototype.filter.call(forms, function (form) {
			form.addEventListener('submit', function (event) {
				if (form.checkValidity() === false) {
					event.preventDefault();
					event.stopPropagation();
				}
				form.classList.add('was-validated');
			}, false);
		});
	}

	function setTime() {


		$(".time-tick").each(function () {
			$(this).on("input", function (e) {
				console.log("input")
				let af = new Date()
				let after = af.getTime()
				let use = after - start
				let name = $(e.delegateTarget).attr("name")
				if (lastName !== name) {
					console.log(name + " use: ", use)
					$("#" + name + "_use").val(use)
					start = (new Date()).getTime()
					lastName = name
				}

			})
		})
	}

	$(function () {
		doShowAfter()
		doVolidate()
		setTime()
	})

</script>
</body>
</html>