class HomeController < ApplicationController
  def index
    item = RssItem
    item = item.where("title like ?", "%"+params[:title] + "%") unless params[:title].blank?
    item = item.where("source like ?", "%"+params[:source]+"%") unless params[:source].blank?

    @data = item.order("publish_date desc")
  end
end
