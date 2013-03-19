$(document).ready(function(){
	prettyPrint();

	$("#comment_form").submit(function(){
		if($("input[name='contetnum']").val()==0){
			alert("请输入回复内容！");
			return false;
		}
		return true;
	})
});


KindEditor.ready(function(K) {
				K.each({ 
					'plug-align' : {
						name : '对齐方式',
						method : {
							'justifyleft' : '左对齐',
							'justifycenter' : '居中对齐',
							'justifyright' : '右对齐'
						}
					},
					'plug-order' : {
						name : '编号',
						method : {
							'insertorderedlist' : '数字编号',
							'insertunorderedlist' : '项目编号'
						}
					},
					'plug-indent' : {
						name : '缩进',
						method : {
							'indent' : '向右缩进',
							'outdent' : '向左缩进'
						}
					}
				},function( pluginName, pluginData ){
					var lang = {};
					lang[pluginName] = pluginData.name;
					KindEditor.lang( lang );
					KindEditor.plugin( pluginName, function(K) {
						var self = this;
						self.clickToolbar( pluginName, function() {
							var menu = self.createMenu({
									name : pluginName,
									width : pluginData.width || 100
								});
							K.each( pluginData.method, function( i, v ){
								menu.addItem({
									title : v,
									checked : false,
									iconClass : pluginName+'-'+i,
									click : function() {
										self.exec(i).hideMenu();
									}
								});
							})
						});
					});
				});
				K.create('#contentqq', {
					themeType : 'sample',
					items : [
						'link'
					],
					afterChange : function() {
						K("input[name='contetnum']").val(this.count('text'));
					}

				});
});