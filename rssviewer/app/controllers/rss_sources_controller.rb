class RssSourcesController < ApplicationController
  def index
    @data = RssSource.all
  end

  def create
    source = RssSource.new
    source.url = params['rss_source']['url']
    if source.save
      RssJob.perform_later
      redirect_to rss_sources_path
    else
      render :new, status: :unprocessable_entity
    end
  end

  def destroy
    @source = RssSource.find(params[:id])
    if @source.delete
      RssJob.perform_later
      redirect_to rss_sources_path
    end
  end

  def new
    @source = RssSource.new 
  end

  def show
    @source = RssSource.find(params[:id])
  end
end
