<?php

declare(strict_types=1);

namespace App\Command;

use App\Service\GisSpider\Review\Review;
use App\Service\GisSpider\Review\ReviewSpider;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Attribute\AsCommand;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;

#[AsCommand(name: 'parse:site:data')]
class ParseSiteData extends Command
{
    public function __construct(
        private ReviewSpider $reviewSpider,
    )
    {
        parent::__construct();
    }

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        $rev = new Review(70000001055902903);

        $this->reviewSpider->handle(1);

        //$output->writeln($data);
        return Command::SUCCESS;
    }
}