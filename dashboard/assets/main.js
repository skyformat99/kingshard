var KS = {};

KS.Configure = {
	load: function(){
		$('#sidebar .active').removeClass('active');
		$('#sidebar .configure').parent().addClass('active');

		$('#content').html($('#configure').html());

		this.global();
		this.nodeState();
		this.schema();
	},
	global: function() {
		var globalConfig = $('#content').find('.global-config');

		var url = "/api/configure/global";

		$.getJSON(url, function(data){
			if (data.code == 0) {
				data = data.data;

				globalConfig.find('.badge').text(data.length);

				if (data.length == 0) {
					globalConfig.find('table tbody').html('<tr><td colspan="2">no global configure!</td></tr>');
					return;
				}

				var configHtml = '';
				for (var k in data) {
					var row = data[k];
					configHtml += "<tr>"+
							"<td>"+row.Key+"</td>"+
							"<td>"+row.Value+"</td>"+
						"</tr>";
				}
				globalConfig.find('table tbody').html(configHtml);

			} else {
				alert("server error:" + data.msg);
			}
		});
	},
	nodeState: function() {
		var nodeState = $('#content').find('.node-state');

		var url = "/api/configure/node_state";

		$.getJSON(url, function(data){
			if (data.code == 0) {
				data = data.data;

				nodeState.find('.badge').text(data.length);

				if (data.length == 0) {
					nodeState.find('table tbody').html('<tr><td colspan="7">no node state!</td></tr>');
					return;
				}

				var configHtml = '';
				for (var k in data) {
					var row = data[k];
					configHtml += "<tr>"+
							"<td>"+row.Node+"</td>"+
							"<td>"+row.Address+"</td>"+
							"<td>"+row.Type+"</td>"+
							"<td>"+row.State+"</td>"+
							"<td>"+row.LastPing+"</td>"+
							"<td>"+row.MaxIdleConn+"</td>"+
							"<td>"+row.IdleConn+"</td>"+
						"</tr>";
				}
				nodeState.find('table tbody').html(configHtml);

			} else {
				alert("server error:" + data.msg);
			}
		});
	},
	schema: function() {
		var schema = $('#content').find('.schema');

		var url = "/api/configure/schema";

		$.getJSON(url, function(data){
			if (data.code == 0) {
				data = data.data;

				schema.find('.badge').text(data.length);

				if (data.length == 0) {
					schema.find('table tbody').html('<tr><td colspan="7">no schema!</td></tr>');
					return;
				}

				var configHtml = '';
				for (var k in data) {
					var row = data[k];
					configHtml += "<tr>"+
							"<td>"+row.DB+"</td>"+
							"<td>"+row.Table+"</td>"+
							"<td>"+row.Type+"</td>"+
							"<td>"+row.Key+"</td>"+
							"<td>"+row.Nodes_List+"</td>"+
							"<td>"+row.Locations+"</td>"+
							"<td>"+row.TableRowLimit+"</td>"+
						"</tr>";
				}
				schema.find('table tbody').html(configHtml);

			} else {
				alert("server error:" + data.msg);
			}
		});
	}

};

KS.DbManage = {
	load: function() {
		$('#sidebar .active').removeClass('active');
		$('#sidebar .db-manage').parent().addClass('active');

		$('#content').html($('#db-manage').html());

		this.nodeList();
	},
	nodeList: function() {
		var nodeList = $('#content').find('.node-list');

		var url = "/api/configure/node_state";

		$.getJSON(url, function(data){
			if (data.code == 0) {
				data = data.data;

				nodeList.find('.badge').text(data.length);

				if (data.length == 0) {
					nodeList.find('table tbody').html('<tr><td colspan="6">no nodes!</td></tr>');
					return;
				}

				var configHtml = '';
				for (var k in data) {
					var row = data[k],
						opTxt = '';
					if (row.State == "up") {
						opTxt = '<button class="btn btn-mini btn-warning">Offline</button>';
					} else {
						opTxt = '<button class="btn btn-mini btn-success">Online</button>';
					}

					configHtml += "<tr>"+
							"<td>"+row.Node+"</td>"+
							"<td>"+row.Address+"</td>"+
							"<td>"+row.Type+"</td>"+
							"<td class='state'>"+row.State+"</td>"+
							"<td>"+row.LastPing+"</td>"+
							"<td>"+opTxt+"</td>"+
						"</tr>";
				}
				nodeList.find('table tbody').html(configHtml);

			} else {
				alert("server error:" + data.msg);
			}
		});
	},
	doManage: function(target, opt) {
		var that = this;
		var $tds = $(target).parents('tr').children();

		var params = {
			'opt': opt,
			'node': $($tds.get(0)).text(),	// nodeName
			'k': $($tds.get(2)).text(),		// type
			'v': $($tds.get(1)).text()		// address
		};

		var url = '/api/db/manage';

		$.post(url, params, function(data) {
			data = $.parseJSON(data)
			if (data.code == 0) {
				// reload
				that.load();
			} else {
				alert("server error:" + data.msg);
			}
		});
	}
};


$(function() {
	
	switch (location.pathname) {
	case '/configure':
		KS.Configure.load();
		break;
	case '/dbmanage':
		KS.DbManage.load();
		break;
	default:
		KS.Configure.load();
	}

	$('#sidebar .configure').on('click', function(evt){
		evt.preventDefault();

		history.pushState({}, 'Configure', location.origin+"/configure")

		KS.Configure.load();
	});

	$('#sidebar .db-manage').on('click', function(evt){
		evt.preventDefault();

		history.pushState({}, 'Configure', location.origin+"/dbmanage")

		KS.DbManage.load();
	});

	$('#content').on('click', '.node-list table .btn', function(){
		var opt = 'down';
		if ($(this).parents('tr').children('.state').text() == "down") {
			opt = "up"
		}
		KS.DbManage.doManage(this, opt);
	});
});